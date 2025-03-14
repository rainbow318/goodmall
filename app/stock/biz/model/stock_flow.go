package model

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	dalredis "github.com/suutest/app/stock/biz/dal/redis"
	"gorm.io/gorm"
)

const (
	CancelStatus           = "canceled"
	SuccessStatus          = "success"
	OccupiedExpireDuration = 15 * time.Minute // 占用库存的时长
	FlowLifetime           = 30 * time.Minute // 库存流水的生命周期
)

type StockFlow struct {
	gorm.Model
	FlowId    string    `redis:"flow_id" gorm:"type:varchar(50)"`
	ProductId uint32    `redis:"prodcut_id" gorm:"type:int(11)"`
	Quantity  uint32    `redis:"quantity"  gorm:"type:int(11)"`
	OrderId   string    `redis:"order_id"  gorm:"type:varchar(50)"`
	Status    string    `redis:"status"  gorm:"type:varchar(10)"`
	CreatedAt time.Time `redis:"created_at" gorm:"type:datetime"`
}

func (StockFlow) TableName() string {
	return "stock_flow"
}

func StructToRedisMap(flow StockFlow) map[string]interface{} {
	return map[string]interface{}{
		"flow_id":    flow.FlowId,
		"product_id": flow.ProductId,
		"quantity":   flow.Quantity,
		"order_id":   flow.OrderId,
		"status":     flow.Status,
		"created_at": flow.CreatedAt.Format(time.RFC3339),
	}
}

func GenerateFlowId(order_id string) string {
	return order_id + time.Now().Format("20060102150405")
}

// 更新库存
func UpdateStock(ctx context.Context, productId uint32, realDelta int, preDelta int) error {
	realKey := fmt.Sprintf("stock:real:%d", productId)
	preKey := fmt.Sprintf("stock:pre:%d", productId)

	pipe := dalredis.RedisClient.Pipeline()
	// 更新实际库存
	if realDelta != 0 {
		pipe.IncrBy(ctx, realKey, int64(realDelta))
	}

	// 更新预占库存
	if preDelta != 0 {
		pipe.IncrBy(ctx, preKey, int64(preDelta))
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		fmt.Printf("update stock error:%s", err.Error())
		klog.Errorf("update stock error:%s", err.Error())
		return err
	}
	return nil
}

func SaveFlowsInDB(ctx context.Context, db *gorm.DB, flows []*StockFlow) error {
	return db.WithContext(ctx).Model(&StockFlow{}).Create(flows).Error
}
