package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/suutest/app/stock/biz/model"
	stock "github.com/suutest/rpc_gen/kitex_gen/stock"
)

type RecoverStocksService struct {
	ctx context.Context
} // NewRecoverStocksService new RecoverStocksService
func NewRecoverStocksService(ctx context.Context) *RecoverStocksService {
	return &RecoverStocksService{ctx: ctx}
}

// Run create note info
func (s *RecoverStocksService) Run(req *stock.RecoverStocksReq) (resp *stock.RecoverStocksResp, err error) {
	for _, sto := range req.Stocks {
		err = model.UpdateStock(s.ctx, sto.ProductId, int(sto.Quantity), -int(sto.Quantity))
		if err != nil {
			klog.Errorf("update stock failed:%s", err.Error())
			continue
		}
	}
	return
}
