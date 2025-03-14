package sync

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/suutest/app/stock/biz/dal/mysql"
	"github.com/suutest/app/stock/biz/dal/redis"
	"github.com/suutest/app/stock/biz/model"
	"gorm.io/gorm"
)

func syncAllStockToRedis(ctx context.Context) {
	stocks, err := model.GetAllStocks(ctx, mysql.DB)
	if err != nil {
		klog.Error("sync all stock fail:", err)
		panic(err)
	}
	for _, s := range stocks {
		key := fmt.Sprintf("stock:real:%d", s.ProductId)
		model.ProductIdList = append(model.ProductIdList, s.ProductId)
		redis.RedisClient.Set(ctx, key, int(s.Quantity), 0)
	}
}

func transaction(ctx context.Context, db *gorm.DB, fn func(tx *gorm.DB) error) error {
	tx := db.Begin()
	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func SyncStockCache(ctx context.Context) {
	syncAllStockToRedis(ctx)

	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	go func() {
		for range ticker.C {
			if err := transaction(ctx, mysql.DB, func(tx *gorm.DB) error {
				syncAllStockToRedis(ctx)
				return nil
			}); err != nil {
				klog.Errorf("sync stock fail:", err)
			}
		}
	}()
}
