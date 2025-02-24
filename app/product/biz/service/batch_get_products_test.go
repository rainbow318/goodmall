package service

import (
	"context"
	"testing"
	product "github.com/suutest/rpc_gen/kitex_gen/product"
)

func TestBatchGetProducts_Run(t *testing.T) {
	ctx := context.Background()
	s := NewBatchGetProductsService(ctx)
	// init req and assert value

	req := &product.BatchGetProductsReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
