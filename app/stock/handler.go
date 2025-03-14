package main

import (
	"context"

	"github.com/suutest/app/stock/biz/service"
	stock "github.com/suutest/rpc_gen/kitex_gen/stock"
)

// StockServiceImpl implements the last service interface defined in the IDL.
type StockServiceImpl struct{}

// BatchGetStocks implements the StockServiceImpl interface.
func (s *StockServiceImpl) BatchGetStocks(ctx context.Context, req *stock.GetStocksReq) (resp *stock.GetStocksResp, err error) {
	resp, err = service.NewBatchGetStocksService(ctx).Run(req)

	return resp, err
}

// DeductStocks implements the StockServiceImpl interface.
func (s *StockServiceImpl) DeductStocks(ctx context.Context, req *stock.DeductStocksReq) (resp *stock.DeductStocksResp, err error) {
	resp, err = service.NewDeductStocksService(ctx).Run(req)

	return resp, err
}

// OccupyStocks implements the StockServiceImpl interface.
func (s *StockServiceImpl) OccupyStocks(ctx context.Context, req *stock.OccupyStocksReq) (resp *stock.OccupyStocksResp, err error) {
	resp, err = service.NewOccupyStocksService(ctx).Run(req)

	return resp, err
}

// RecoverStocks implements the StockServiceImpl interface.
func (s *StockServiceImpl) RecoverStocks(ctx context.Context, req *stock.RecoverStocksReq) (resp *stock.RecoverStocksResp, err error) {
	resp, err = service.NewRecoverStocksService(ctx).Run(req)

	return resp, err
}
