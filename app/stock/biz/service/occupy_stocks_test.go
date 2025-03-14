package service

import (
	"context"
	"testing"

	stock "github.com/suutest/rpc_gen/kitex_gen/stock"
)

func TestOccupyStocks_Run(t *testing.T) {
	ctx := context.Background()
	s := NewOccupyStocksService(ctx)
	// init req and assert value
	var stocks []*stock.Stock
	stocks = append(stocks, &stock.Stock{
		ProductId: 1,
		Quantity:  1,
	})
	stocks = append(stocks, &stock.Stock{
		ProductId: 2,
		Quantity:  1,
	})
	req := &stock.OccupyStocksReq{
		Stocks:  stocks,
		OrderId: "test1",
	}

	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
