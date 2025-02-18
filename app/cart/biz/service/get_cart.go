package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/suutest/app/cart/biz/dal/mysql"
	"github.com/suutest/app/cart/biz/model"
	cart "github.com/suutest/rpc_gen/kitex_gen/cart"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	if req.UserId == 0 {
		return nil, kerrors.NewBizStatusError(50002, "user id is required")
	}
	itemsList, err := model.GetCartByUserId(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50002, err.Error())
	}
	var items []*cart.CartItem
	for _, v := range itemsList {
		items = append(items, &cart.CartItem{
			ProductId: v.ProductId,
			Quantity:  v.Qty,
		})
	}
	resp = &cart.GetCartResp{
		Items: items,
	}
	return resp, nil
}
