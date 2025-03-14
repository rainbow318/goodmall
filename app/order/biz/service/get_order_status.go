package service

import (
	"context"

	"github.com/suutest/app/order/biz/dal/mysql"
	"github.com/suutest/app/order/biz/model"
	order "github.com/suutest/rpc_gen/kitex_gen/order"
)

type GetOrderStatusService struct {
	ctx context.Context
} // NewGetOrderStatusService new GetOrderStatusService
func NewGetOrderStatusService(ctx context.Context) *GetOrderStatusService {
	return &GetOrderStatusService{ctx: ctx}
}

// Run create note info
func (s *GetOrderStatusService) Run(req *order.GetOrderStatusReq) (resp *order.GetOrderStatusResp, err error) {
	// Finish your business logic.
	status, err := model.GetOrderStatus(s.ctx, mysql.DB, req.OrderId)
	if err != nil {
		return nil, err
	}
	return &order.GetOrderStatusResp{Status: status}, nil
}
