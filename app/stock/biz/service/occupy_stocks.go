package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/suutest/app/stock/biz/dal/redis"
	"github.com/suutest/app/stock/biz/model"
	stock "github.com/suutest/rpc_gen/kitex_gen/stock"
)

type OccupyStocksService struct {
	ctx context.Context
} // NewOccupyStocksService new OccupyStocksService
func NewOccupyStocksService(ctx context.Context) *OccupyStocksService {
	return &OccupyStocksService{ctx: ctx}
}

// Run create note info
func (s *OccupyStocksService) Run(req *stock.OccupyStocksReq) (resp *stock.OccupyStocksResp, err error) {
	// 检查可用库存
	// stock, err := model.GetStockById(s.ctx, mysql.DB, req.Stock.ProductId)
	for _, sto := range req.Stocks {
		item, err := model.GetStockByIdFromRedis(s.ctx, redis.RedisClient, sto.ProductId)
		if err != nil {
			klog.Errorf("get stock failed:%s", err.Error())
			fmt.Printf("get stock failed:%s", err.Error())
			return nil, err
		}
		if uint32(item.Quantity) < sto.Quantity {
			fmt.Printf("productId=%d, stock not enough", sto.ProductId)
			return nil, fmt.Errorf("productId=%d, stock not enough", sto.ProductId)
		}
	}

	for _, sto := range req.Stocks {
		// 更新库存（此时更新的只是预扣的库存，预扣数量增加，减少可用库存）
		err = model.UpdateStock(s.ctx, sto.ProductId, -(int(sto.Quantity)), int(sto.Quantity))
		if err != nil {
			fmt.Printf("update stock failed:%s", err.Error())
			klog.Errorf("update stock failed:%s", err.Error())
			return nil, err
		}
	}
	return &stock.OccupyStocksResp{
		Success: true,
	}, nil
}
