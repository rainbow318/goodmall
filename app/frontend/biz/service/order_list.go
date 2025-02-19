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
	for _, v := range orderResp.Orders {
		var (
			total float32
			items []types.OrderItem
		)
		// TODO：生产中不会直接在循环中调用RPC，而是在外面组装prodcutID然后批量获取product信息，再去组装map
		for _, i := range v.Items {
			total += i.Cost

			// 商品名，图片路径需要通过rpc调用
			productResp, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductReq{
				Id: i.Item.ProductId,
			})
			if err != nil {
				return nil, err
			}
			if productResp == nil || productResp.Product == nil {
				continue
			}

			items = append(items, types.OrderItem{
				ProductName: productResp.Product.Name,
				Picture:     productResp.Product.Picture,
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
