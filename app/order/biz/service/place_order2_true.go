package service

import (
	"context"

	"github.com/suutest/app/order/biz/dal/mysql"
	"github.com/suutest/app/order/biz/model"
	order "github.com/suutest/rpc_gen/kitex_gen/order"
)

type PlaceOrder2TrueService struct {
	ctx context.Context
} // NewPlaceOrder2TrueService new PlaceOrder2TrueService
func NewPlaceOrder2TrueService(ctx context.Context) *PlaceOrder2TrueService {
	return &PlaceOrder2TrueService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrder2TrueService) Run(req *order.PlaceOrder2TrueReq) (resp *order.PlaceOrder2TrueResp, err error) {
	err = model.SetIsCharged2True(s.ctx, mysql.DB, req.OrderId)
	if err != nil {
		return nil, err
	}
	return &order.PlaceOrder2TrueResp{}, nil
}
