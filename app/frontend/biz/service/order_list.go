package service

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	common "github.com/suutest/app/frontend/hertz_gen/frontend/common"
	"github.com/suutest/app/frontend/infra/rpc"
	"github.com/suutest/app/frontend/types"
	frontendUtils "github.com/suutest/app/frontend/utils"
	"github.com/suutest/rpc_gen/kitex_gen/order"
	"github.com/suutest/rpc_gen/kitex_gen/product"
)

type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run(req *common.Empty) (resp map[string]any, err error) {
	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	orderResp, err := rpc.OrderClient.ListOrder(h.Context, &order.ListOrderReq{
		UserId: uint32(userId),
	})
	if err != nil {
		return nil, err
	}
	var orderList []types.Order

	var ids []uint32
	// 获取需要查询的商品的id
	for _, v := range orderResp.Orders {
		for _, i := range v.Items {
			ids = append(ids, i.Item.ProductId)
		}
	}

	products, err := rpc.ProductClient.BatchGetProducts(h.Context, &product.BatchGetProductsReq{
		Ids: ids,
	})

	for _, v := range orderResp.Orders {
		var (
			total float32
			items []types.OrderItem
		)
		for _, i := range v.Items {
			total += i.Cost
			p := products.Products[uint32(i.Item.ProductId)]
			items = append(items, types.OrderItem{
				ProductId:   i.Item.ProductId,
				ProductName: p.Name,
				Picture:     p.Picture,
				Qty:         i.Item.Quantity,
				Cost:        i.Cost,
			})
		}
		created := time.Unix(int64(v.CreatedAt), 0)
		orderList = append(orderList, types.Order{
			OrderId:     v.OrderId,
			CreatedDate: created.Format("2006-01-02 15:04:05"),
			Cost:        total,
			Items:       items,
		})
	}
	return utils.H{
		"title":  "Order",
		"orders": orderList,
	}, nil
}
