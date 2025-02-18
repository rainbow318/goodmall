package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/suutest/app/cart/biz/dal/mysql"
	"github.com/suutest/app/cart/biz/model"
	"github.com/suutest/app/cart/infra/rpc"
	cart "github.com/suutest/rpc_gen/kitex_gen/cart"
	rpcproduct "github.com/suutest/rpc_gen/kitex_gen/product"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	productResp, err := rpc.ProductClient.GetProduct(s.ctx, &rpcproduct.GetProductReq{
		Id: uint32(req.Item.ProductId),
	})
	if err != nil {
		return nil, err
	}
	if productResp == nil || req.Item.ProductId == 0 {
		return nil, kerrors.NewBizStatusError(40000, "product not found") // 40000是随便定义的错误码
	}
	cartItem := &model.Cart{
		UserId:    req.UserId,
		ProductId: req.Item.ProductId,
		Qty:       req.Item.Quantity,
	}
	err = model.AddItem(s.ctx, mysql.DB, cartItem)
	if err != nil {
		return nil, kerrors.NewBizStatusError(400003, err.Error())
	}
	return &cart.AddItemResp{}, nil
}
