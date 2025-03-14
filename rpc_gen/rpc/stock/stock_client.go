package stock

import (
	"context"
	stock "github.com/suutest/rpc_gen/kitex_gen/stock"

	"github.com/suutest/rpc_gen/kitex_gen/stock/stockservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() stockservice.Client
	Service() string
	BatchGetStocks(ctx context.Context, Req *stock.GetStocksReq, callOptions ...callopt.Option) (r *stock.GetStocksResp, err error)
	DeductStocks(ctx context.Context, Req *stock.DeductStocksReq, callOptions ...callopt.Option) (r *stock.DeductStocksResp, err error)
	OccupyStocks(ctx context.Context, Req *stock.OccupyStocksReq, callOptions ...callopt.Option) (r *stock.OccupyStocksResp, err error)
	RecoverStocks(ctx context.Context, Req *stock.RecoverStocksReq, callOptions ...callopt.Option) (r *stock.RecoverStocksResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := stockservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient stockservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() stockservice.Client {
	return c.kitexClient
}

func (c *clientImpl) BatchGetStocks(ctx context.Context, Req *stock.GetStocksReq, callOptions ...callopt.Option) (r *stock.GetStocksResp, err error) {
	return c.kitexClient.BatchGetStocks(ctx, Req, callOptions...)
}

func (c *clientImpl) DeductStocks(ctx context.Context, Req *stock.DeductStocksReq, callOptions ...callopt.Option) (r *stock.DeductStocksResp, err error) {
	return c.kitexClient.DeductStocks(ctx, Req, callOptions...)
}

func (c *clientImpl) OccupyStocks(ctx context.Context, Req *stock.OccupyStocksReq, callOptions ...callopt.Option) (r *stock.OccupyStocksResp, err error) {
	return c.kitexClient.OccupyStocks(ctx, Req, callOptions...)
}

func (c *clientImpl) RecoverStocks(ctx context.Context, Req *stock.RecoverStocksReq, callOptions ...callopt.Option) (r *stock.RecoverStocksResp, err error) {
	return c.kitexClient.RecoverStocks(ctx, Req, callOptions...)
}
