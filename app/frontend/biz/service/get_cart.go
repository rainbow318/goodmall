package service

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	common "github.com/suutest/app/frontend/hertz_gen/frontend/common"
	"github.com/suutest/app/frontend/infra/rpc"
	frontendUtils "github.com/suutest/app/frontend/utils"
	"github.com/suutest/rpc_gen/kitex_gen/cart"
	"github.com/suutest/rpc_gen/kitex_gen/product"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *common.Empty) (resp map[string]any, err error) {
	cartResp, err := rpc.CartClient.GetCart(h.Context, &cart.GetCartReq{
		UserId: uint32(frontendUtils.GetUserIdFromCtx(h.Context)),
	})
	if err != nil {
		return nil, err
	}
	var items []map[string]string
	var total float64 // 用来记录购物车所有商品的总价
	for _, item := range cartResp.Items {
		// 得到的item只包含产品id和数量，但是渲染前端页面时只有产品id和数量是不够的。所以还要用rpc调用product
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductReq{
			Id: item.ProductId,
		})
		if err != nil {
			continue // 这里是简单处理：拿商品信息有错误的就跳过
		}
		p := productResp.Product
		items = append(items, map[string]string{
			"Name":        p.Name,
			"Description": p.Description,
			"Price":       strconv.FormatFloat(float64(p.Price), 'f', 2, 64), // 保留两位有效小数
			"Picture":     p.Picture,
			"Qty":         strconv.Itoa(int(item.Quantity)),
		})
		total += float64(item.Quantity) * float64(p.Price)
	}
	return utils.H{
		"title": "Cart",
		"items": items,
		"total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
