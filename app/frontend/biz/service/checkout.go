package service

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	common "github.com/suutest/app/frontend/hertz_gen/frontend/common"
	"github.com/suutest/app/frontend/infra/rpc"
	frontendUtils "github.com/suutest/app/frontend/utils"
	rpccart "github.com/suutest/rpc_gen/kitex_gen/cart"
	rpcproduct "github.com/suutest/rpc_gen/kitex_gen/product"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *common.Empty) (resp map[string]any, err error) {
	// 展示给用户一个需要结算的商品列表
	var items []map[string]string
	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	carts, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{UserId: uint32(userId)})
	if err != nil {
		return nil, err
	}
	var total float32
	for _, v := range carts.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{
			Id: v.ProductId,
		})
		if err != nil {
			return nil, err
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product
		items = append(items, map[string]string{
			"Name":        p.Name,
			"Description": p.Description,
			"Price":       strconv.FormatFloat(float64(p.Price), 'f', 2, 64), // 保留两位有效小数
			"Picture":     p.Picture,
			"Qty":         strconv.Itoa(int(v.Quantity)),
		})
		total += p.Price * float32(v.Quantity)
	}
	return utils.H{
		"title": "Checkout",
		"items": items,
		"total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
