package service

import (
	"context"
	"testing"
	order "github.com/suutest/rpc_gen/kitex_gen/order"
)

func TestPlaceOrder2True_Run(t *testing.T) {
	ctx := context.Background()
	s := NewPlaceOrder2TrueService(ctx)
	// init req and assert value

	req := &order.PlaceOrder2TrueReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
