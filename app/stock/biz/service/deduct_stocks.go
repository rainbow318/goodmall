package service

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/suutest/app/stock/biz/dal/mysql"
	"github.com/suutest/app/stock/biz/model"
	stock "github.com/suutest/rpc_gen/kitex_gen/stock"
)

type DeductStocksService struct {
	ctx context.Context
} // NewDeductStocksService new DeductStocksService
func NewDeductStocksService(ctx context.Context) *DeductStocksService {
	return &DeductStocksService{ctx: ctx}
}

// Run create note info
func (s *DeductStocksService) Run(req *stock.DeductStocksReq) (resp *stock.DeductStocksResp, err error) {
	var stocks []*model.Stock
	for _, v := range req.Stocks {
		stocks = append(stocks, &model.Stock{
			ProductId: v.ProductId,
			Quantity:  int(v.Quantity),
		})
	}
	ok, err := model.DeductStocks(s.ctx, mysql.DB, stocks, req.OrderId)

	// 记录库存扣减的流水
	var flows []*model.StockFlow
	flowId := model.GenerateFlowId(req.OrderId)
	creatAt := time.Now()

	if err != nil || !ok {
		for _, v := range req.Stocks {
			flows = append(flows, &model.StockFlow{
				FlowId:    flowId,
				ProductId: v.ProductId,
				Quantity:  v.Quantity,
				OrderId:   req.OrderId,
				Status:    model.CancelStatus,
				CreatedAt: creatAt,
			})
		}
		if e := model.SaveFlowsInDB(s.ctx, mysql.DB, flows); e != nil {
			klog.Errorf("save stock flows failed:%s", e.Error())
		}
		return nil, err
	} else {
		for _, v := range req.Stocks {
			flows = append(flows, &model.StockFlow{
				FlowId:    flowId,
				ProductId: v.ProductId,
				Quantity:  v.Quantity,
				OrderId:   req.OrderId,
				Status:    model.SuccessStatus,
				CreatedAt: creatAt,
			})
		}
		if e := model.SaveFlowsInDB(s.ctx, mysql.DB, flows); e != nil {
			klog.Errorf("save stock flows failed:%s", e.Error())
		}
	}

	return &stock.DeductStocksResp{
		Success: ok,
	}, nil
}
