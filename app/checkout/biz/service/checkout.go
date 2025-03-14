package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/nats-io/nats.go"
	"github.com/suutest/app/checkout/infra/mq"
	"github.com/suutest/app/checkout/infra/rpc"
	"github.com/suutest/rpc_gen/kitex_gen/cart"
	checkout "github.com/suutest/rpc_gen/kitex_gen/checkout"
	rpcemail "github.com/suutest/rpc_gen/kitex_gen/email"
	"github.com/suutest/rpc_gen/kitex_gen/order"
	"github.com/suutest/rpc_gen/kitex_gen/payment"
	"github.com/suutest/rpc_gen/kitex_gen/product"
	"github.com/suutest/rpc_gen/kitex_gen/stock"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/protobuf/proto"
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

	var (
		total float32
		oi    []*order.OrderItem
	)
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
		cost := float32(cartItem.Quantity) * productResp.Product.Price
		total += cost
		oi = append(oi, &order.OrderItem{
			Item: &cart.CartItem{
				ProductId: cartItem.ProductId,
				Quantity:  cartItem.Quantity,
			},
			Cost: cost,
		})
	}

	var orderId string

	orderResp, err := rpc.OrderClient.PlaceOrder(s.ctx, &order.PlaceOrderReq{
		UserId: req.UserId,
		Email:  req.Email,
		Address: &order.Address{
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			State:         req.Address.State,
			Country:       req.Address.City,
			ZipCode:       req.Address.ZipCode,
		},
		Items:     oi,
		IsCharged: 0,
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5004002, err.Error())
	}

	if orderResp != nil && orderResp.Order != nil {
		orderId = orderResp.Order.OrderId
	}

	// 预扣减库存
	var stocks []*stock.Stock
	for _, cartItem := range cartResult.Items {
		stocks = append(stocks, &stock.Stock{
			ProductId: cartItem.ProductId,
			Quantity:  cartItem.Quantity,
		})
	}
	occupyResp, err := rpc.StockClient.OccupyStocks(s.ctx, &stock.OccupyStocksReq{
		Stocks:  stocks,
		OrderId: orderId,
	})
	if err != nil || !occupyResp.Success {
		klog.Errorf("occupy stock fail")
		return nil, err
	}

	payReq := &payment.ChargeReq{
		Amount:     total,
		CreditCard: req.CreditCard,
		OrderId:    orderId,
		UserId:     req.UserId,
	}
	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	// 如果支付失败，则恢复库存
	if err != nil {
		klog.Errorf("charge fail:%s", err.Error())
		// 恢复库存
		recoverResp, err := rpc.StockClient.OccupyStocks(s.ctx, &stock.OccupyStocksReq{
			Stocks:  stocks,
			OrderId: orderId,
		})
		if err != nil || !recoverResp.Success {
			klog.Errorf("recover stock fail:%s", err.Error())
			return nil, err
		}
		return nil, err
	}
	// 真正扣库存
	deductStockResp, err := rpc.StockClient.DeductStocks(s.ctx, &stock.DeductStocksReq{
		Stocks:  stocks,
		OrderId: orderId,
	})
	if err != nil || !deductStockResp.Success {
		klog.Errorf("deduct stock fail:%s", err.Error())
		return nil, err
	}

	// 更新订单状态：order表中is_charged字段设为1
	_, err = rpc.OrderClient.PlaceOrder2True(s.ctx, &order.PlaceOrder2TrueReq{OrderId: orderId})
	if err != nil {
		return nil, err
	}

	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{
		UserId: req.UserId,
	})
	if err != nil {
		klog.Error(err.Error())
	}
	// 消息中间件 生产者
	data, _ := proto.Marshal(&rpcemail.EmailReq{
		From:        "from@example.com",
		To:          req.Email,
		ContentType: "text/plain",
		Subject:     "GoodMall: 订单信息",
		Content:     "您刚在GoodMall新建了一个订单",
	})
	msg := &nats.Msg{ // 生产信息时需要将链路信息载入到消息的header部分
		Subject: "email",
		Data:    data,
		Header:  make(nats.Header),
	}
	otel.GetTextMapPropagator().Inject(s.ctx, propagation.HeaderCarrier(msg.Header))
	err = mq.Nc.PublishMsg(msg)
	if err != nil {
		klog.Error(err)
	}

	klog.Info(paymentResult)

	resp = &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResult.TransactionId, // 校验id
	}
	return
}
