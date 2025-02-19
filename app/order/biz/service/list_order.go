package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/suutest/app/order/biz/dal/mysql"
	"github.com/suutest/app/order/biz/model"
	"github.com/suutest/rpc_gen/kitex_gen/cart"
	order "github.com/suutest/rpc_gen/kitex_gen/order"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	list, err := model.ListOrder(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(600002, err.Error())
	}
	var orders []*order.Order
	for _, v := range list {
		var items []*order.OrderItem
		for _, oi := range v.OrderItems {
			items = append(items, &order.OrderItem{
				Item: &cart.CartItem{
					ProductId: oi.ProductId,
					Quantity:  oi.Quantity,
				},
				Cost: oi.Cost,
			})
		}
		orders = append(orders, &order.Order{
			OrderId:      v.OrderId,
			UserId:       v.UserId,
			UserCurrency: v.UserCurrenty,
			Email:        v.Consignee.Email,
			Address: &order.Address{
				StreetAddress: v.Consignee.Street,
				City:          v.Consignee.City,
				State:         v.Consignee.State,
				Country:       v.Consignee.Country,
				ZipCode:       v.Consignee.ZipCode,
				Phone:         v.Consignee.Phone,
			},
			Items:     items,
			CreatedAt: int32(v.CreatedAt.Unix()),
		})
	}
	resp = &order.ListOrderResp{
		Orders: orders,
	}
	return resp, nil
}
