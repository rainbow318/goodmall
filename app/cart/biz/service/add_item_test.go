package service

import (
	"context"
	"reflect"
	"testing"

	cart "github.com/suutest/rpc_gen/kitex_gen/cart"
)

func TestAddItem_Run(t *testing.T) {
	ctx := context.Background()
	s := NewAddItemService(ctx)
	// init req and assert value

	req := &cart.AddItemReq{
		UserId: uint32(1),
		Item: &cart.CartItem{
			ProductId: uint32(2),
			Quantity:  uint32(3),
		},
	}

	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}

func TestAddItemService_Run(t *testing.T) {
	type fields struct {
		ctx context.Context
	}
	type args struct {
		req *cart.AddItemReq
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *cart.AddItemResp
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AddItemService{
				ctx: tt.fields.ctx,
			}
			gotResp, err := s.Run(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddItemService.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("AddItemService.Run() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
