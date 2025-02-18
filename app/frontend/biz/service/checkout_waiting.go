package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	checkout "github.com/suutest/app/frontend/hertz_gen/frontend/checkout"
	"github.com/suutest/app/frontend/infra/rpc"
	frontendUtils "github.com/suutest/app/frontend/utils"
	rpccheckout "github.com/suutest/rpc_gen/kitex_gen/checkout"
	rpcpayment "github.com/suutest/rpc_gen/kitex_gen/payment"
)

type CheckoutWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutWaitingService(Context context.Context, RequestContext *app.RequestContext) *CheckoutWaitingService {
	return &CheckoutWaitingService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutWaitingService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	// 这个接口才是真实的checkout接口，需要调用checkout服务来做最终的结算
	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	_, err = rpc.CheckoutClient.Checkout(h.Context, &rpccheckout.CheckoutReq{
		UserId:    uint32(userId),
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Address: &rpccheckout.Address{
			StreetAddress: req.Street,
			City:          req.City,
			State:         req.Province,
			Country:       req.Country,
			ZipCode:       req.Zipcode,
		},
		CreditCard: &rpcpayment.CreditCardInfo{
			CreditCardNumber:          req.CardNumber,
			CreditCardCvv:             req.Cvv,
			CreditCardExpirationYear:  req.ExpirationYear,
			CreditCardExpirationMonth: req.ExpirationMonth,
		},
	})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"title":    "checkout_waiting",
		"redirect": "/checkout/result",
	}, nil
}
