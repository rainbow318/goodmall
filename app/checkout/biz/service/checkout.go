package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"github.com/suutest/app/checkout/infra/rpc"
	"github.com/suutest/rpc_gen/kitex_gen/cart"
	checkout "github.com/suutest/rpc_gen/kitex_gen/checkout"
	"github.com/suutest/rpc_gen/kitex_gen/payment"
	"github.com/suutest/rpc_gen/kitex_gen/product"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005001, err.Error())
	}

	if cartResult == nil || cartResult.Items == nil {
		return nil, kerrors.NewGRPCBizStatusError(5004001, "cart is empty")
	}

	var total float32
	// TODO 这里在循环中使用rpc调用，在真实情况下要避免这样做，因为这样对性能会有很大影响。
	// 应该在for循环外面统一使用rpc获取数据，然后再遍历计算total值
	for _, cartItem := range cartResult.Items {
		productResp, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{
			Id: cartItem.ProductId,
		})
		if err != nil {
			return nil, err
		}
		if productResp.Product == nil {
			continue
		}
		total += float32(cartItem.Quantity) * productResp.Product.Price
	}

	var orderId string
	// TODO 下面应该调用订单服务。这里暂时先使用一个虚拟订单
	u, _ := uuid.NewRandom()
	orderId = u.String()

	payReq := &payment.ChargeReq{
		Amount:     total,
		CreditCard: req.CreditCard,
		OrderId:    orderId,
		UserId:     req.UserId,
	}

	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{
		UserId: req.UserId,
	})
	if err != nil {
		klog.Error(err.Error())
	}

	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		return nil, err
	}

	klog.Info(paymentResult)

	resp = &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResult.TransactionId, // 校验id
	}
	return
}
