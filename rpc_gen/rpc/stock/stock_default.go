package stock

import (
	"context"
	stock "github.com/suutest/rpc_gen/kitex_gen/stock"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func BatchGetStocks(ctx context.Context, req *stock.GetStocksReq, callOptions ...callopt.Option) (resp *stock.GetStocksResp, err error) {
	resp, err = defaultClient.BatchGetStocks(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "BatchGetStocks call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeductStocks(ctx context.Context, req *stock.DeductStocksReq, callOptions ...callopt.Option) (resp *stock.DeductStocksResp, err error) {
	resp, err = defaultClient.DeductStocks(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeductStocks call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func OccupyStocks(ctx context.Context, req *stock.OccupyStocksReq, callOptions ...callopt.Option) (resp *stock.OccupyStocksResp, err error) {
	resp, err = defaultClient.OccupyStocks(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "OccupyStocks call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func RecoverStocks(ctx context.Context, req *stock.RecoverStocksReq, callOptions ...callopt.Option) (resp *stock.RecoverStocksResp, err error) {
	resp, err = defaultClient.RecoverStocks(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "RecoverStocks call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
