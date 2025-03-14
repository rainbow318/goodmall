package service

import (
	"context"
	"testing"
	stock "github.com/suutest/rpc_gen/kitex_gen/stock"
)

func TestDeductStocks_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeductStocksService(ctx)
	// init req and assert value

	req := &stock.DeductStocksReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
