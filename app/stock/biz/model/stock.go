package model

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	ProductIdList []uint32
	maxRetries    = 3
)

type Stock struct {
	gorm.Model
	ProductId uint32 `gorm:"type:int(11)"`
	Quantity  int    `gorm:"type:int(11)"`
	Version   int    `gorm:"type:int(11);default:0"` // 基于version的乐观锁
}

func (Stock) TableName() string {
	return "stock"
}

func BatchGetStocksByIds(ctx context.Context, db *gorm.DB, productIds []uint32) (stocks []*Stock, err error) {
	err = db.WithContext(ctx).Model(&Stock{}).Where("product_id IN ?", productIds).Find(&stocks).Error
	return
}

func GetStockById(ctx context.Context, db *gorm.DB, productId uint32) (s *Stock, err error) {
	err = db.WithContext(ctx).Model(&Stock{}).Where("product_id ==?", productId).First(&s).Error
	return
}

func GetStockByIdFromRedis(ctx context.Context, r *redis.Client, productId uint32) (s *Stock, err error) {
	stockKey := fmt.Sprintf("stock:real:%d", productId)
	stock_real, err := r.Get(ctx, stockKey).Result()
	if err != nil {
		return nil, err
	}
	quantity, _ := strconv.Atoi(stock_real)
	return &Stock{
		ProductId: productId,
		Quantity:  quantity,
	}, nil
}

func GetStockWithVersion(ctx context.Context, db *gorm.DB, productId uint32) (*Stock, error) {
	var s Stock
	err := db.WithContext(ctx).Debug().Model(&Stock{}).Where("product_id=?", productId).First(&s).Error
	return &s, err
}

// 减库存
func DeductStocks(ctx context.Context, db *gorm.DB, stocks []*Stock, orderId string) (bool, error) {
	// 真正扣减库存
	for t := 0; t < maxRetries; t++ {
		tx := db.Begin()
		commitFlag := true
		for _, s := range stocks {
			// 获取version
			currentStock, err := GetStockWithVersion(ctx, tx, s.ProductId)
			fmt.Printf("currentStock:%v", currentStock.Version)
			if err != nil {
				commitFlag = false // 需要重试
				tx.Rollback()
				break
			}
			if currentStock.Quantity < s.Quantity {
				tx.Rollback()
				return false, fmt.Errorf("productId=%d, stock not enough", s.ProductId)
			}
			updateResult := tx.WithContext(ctx).Debug().
				Model(&Stock{}).
				Where("product_id=? and version=?", s.ProductId, currentStock.Version).
				Updates(map[string]interface{}{
					"quantity": gorm.Expr("quantity-?", s.Quantity),
					"version":  gorm.Expr("version+1"),
				})
			// fmt.Printf("updateResult:%d", updateResult.RowsAffected)
			if updateResult.RowsAffected == 0 {
				tx.Rollback()
				commitFlag = false // 需要重试
				break
			}
		}
		if commitFlag {
			err := tx.Commit().Error
			if err == nil {
				return true, nil
			}
		}
		time.Sleep(time.Duration(t*t) * 100 * time.Microsecond) // 指数退避
	}

	return false, fmt.Errorf("deduct stocks failed")
}

func GetAllStocks(ctx context.Context, db *gorm.DB) (stocks []*Stock, err error) {
	err = db.WithContext(ctx).Model(&Stock{}).Find(&stocks).Error
	return
}
