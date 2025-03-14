package service

import (
	"context"
	"testing"
	stock "github.com/suutest/rpc_gen/kitex_gen/stock"
)

func TestRecoverStocks_Run(t *testing.T) {
	ctx := context.Background()
	s := NewRecoverStocksService(ctx)
	// init req and assert value

	req := &stock.RecoverStocksReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
