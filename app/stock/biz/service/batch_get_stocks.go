package service

import (
	"context"

	"github.com/suutest/app/stock/biz/dal/mysql"
	"github.com/suutest/app/stock/biz/model"
	"github.com/suutest/rpc_gen/kitex_gen/stock"
)

type BatchGetStocksService struct {
	ctx context.Context
} // NewBatchGetStocksService new BatchGetStocksService
func NewBatchGetStocksService(ctx context.Context) *BatchGetStocksService {
	return &BatchGetStocksService{ctx: ctx}
}

// Run create note info
func (s *BatchGetStocksService) Run(req *stock.GetStocksReq) (resp *stock.GetStocksResp, err error) {
	stocksResp, err := model.BatchGetStocksByIds(s.ctx, mysql.DB, req.ProductIds)
	tmp := make(map[uint32]*stock.Stock)
	if err != nil {
		return nil, err
	}
	for _, v := range stocksResp {
		tmp[uint32(v.ProductId)] = &stock.Stock{
			ProductId: v.ProductId,
			Quantity:  uint32(v.Quantity),
		}
	}
	return &stock.GetStocksResp{
		Stocks: tmp,
	}, nil
}
