package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	product "github.com/suutest/app/frontend/hertz_gen/frontend/product"
	"github.com/suutest/app/frontend/infra/rpc"
	rpcproduct "github.com/suutest/rpc_gen/kitex_gen/product"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *product.ProductReq) (resp map[string]any, err error) {
	// 调用product服务的方法
	p, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{
		Id: uint32(req.Id),
	})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"item": p.Product,
	}, nil
}
