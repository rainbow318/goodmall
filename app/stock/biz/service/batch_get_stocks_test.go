package service

import (
	"context"
	"testing"
	stock "github.com/suutest/rpc_gen/kitex_gen/stock"
)

func TestBatchGetStocks_Run(t *testing.T) {
	ctx := context.Background()
	s := NewBatchGetStocksService(ctx)
	// init req and assert value

	req := &stock.GetStocksReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
