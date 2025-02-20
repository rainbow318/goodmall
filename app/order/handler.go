package main

import (
	"context"
	order "github.com/suutest/rpc_gen/kitex_gen/order"
	"github.com/suutest/app/order/biz/service"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// PlaceOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) PlaceOrder(ctx context.Context, req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	resp, err = service.NewPlaceOrderService(ctx).Run(req)

	return resp, err
}

// ListOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ListOrder(ctx context.Context, req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	resp, err = service.NewListOrderService(ctx).Run(req)

	return resp, err
}

// PlaceOrder2True implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) PlaceOrder2True(ctx context.Context, req *order.PlaceOrder2TrueReq) (resp *order.PlaceOrder2TrueResp, err error) {
	resp, err = service.NewPlaceOrder2TrueService(ctx).Run(req)

	return resp, err
}
