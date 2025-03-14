package order

import (
	"context"
	order "github.com/suutest/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func PlaceOrder(ctx context.Context, req *order.PlaceOrderReq, callOptions ...callopt.Option) (resp *order.PlaceOrderResp, err error) {
	resp, err = defaultClient.PlaceOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "PlaceOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ListOrder(ctx context.Context, req *order.ListOrderReq, callOptions ...callopt.Option) (resp *order.ListOrderResp, err error) {
	resp, err = defaultClient.ListOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func PlaceOrder2True(ctx context.Context, req *order.PlaceOrder2TrueReq, callOptions ...callopt.Option) (resp *order.PlaceOrder2TrueResp, err error) {
	resp, err = defaultClient.PlaceOrder2True(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "PlaceOrder2True call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetOrderStatus(ctx context.Context, req *order.GetOrderStatusReq, callOptions ...callopt.Option) (resp *order.GetOrderStatusResp, err error) {
	resp, err = defaultClient.GetOrderStatus(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetOrderStatus call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
