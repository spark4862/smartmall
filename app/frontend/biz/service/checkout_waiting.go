package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	checkout "github.com/spark4862/smartmall/app/frontend/hertz_gen/frontend/checkout"
	"github.com/spark4862/smartmall/app/frontend/infra/rpc"
	frontendUtils "github.com/spark4862/smartmall/app/frontend/utils"
	rpcCheckout "github.com/spark4862/smartmall/rpc_gen/kitex_gen/checkout"
	rpcPayment "github.com/spark4862/smartmall/rpc_gen/kitex_gen/payment"
)

type CheckoutWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutWaitingService(Context context.Context, RequestContext *app.RequestContext) *CheckoutWaitingService {
	return &CheckoutWaitingService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutWaitingService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	_, err = rpc.CheckoutClient.Checkout(h.Context, &rpcCheckout.CheckoutReq{
		UserId:    uint32(userId),
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Address: &rpcCheckout.Address{
			Street:  req.Street,
			City:    req.City,
			State:   req.Province,
			Zip:     req.Zipcode,
			Country: req.Country,
		},
		CreditCard: &rpcPayment.CreditCard{
			Id:              req.CardNum,
			Cvv:             req.Cvv,
			ExpirationYear:  req.ExpYear,
			ExpirationMonth: req.ExpMonth,
		},
	})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"title":    "waiting",
		"redirect": "/checkout/result",
	}, nil
}
