package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/nats-io/nats.go"
	"github.com/spark4862/smartmall/app/checkout/infra/mq"
	"github.com/spark4862/smartmall/app/checkout/infra/rpc"
	"github.com/spark4862/smartmall/rpc_gen/kitex_gen/cart"
	checkout "github.com/spark4862/smartmall/rpc_gen/kitex_gen/checkout"
	"github.com/spark4862/smartmall/rpc_gen/kitex_gen/email"
	"github.com/spark4862/smartmall/rpc_gen/kitex_gen/order"
	"github.com/spark4862/smartmall/rpc_gen/kitex_gen/payment"
	"github.com/spark4862/smartmall/rpc_gen/kitex_gen/product"
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
	// Finish your business logic.
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005001, err.Error())
		// 如需额外传递 gRPC Detail，可以使用 NewGRPCBizStatusError 或 NewGRPCBizStatusErrorWithExtra 来构造异常：
	}
	if cartResult == nil || cartResult.Items == nil {
		return nil, kerrors.NewGRPCBizStatusError(5004001, "cart is empty")
	}

	var total float32
	var items []*order.OrderItem

	// 这里使用for rpc，生产环境中不能这样用，应该提供批量获取接口
	for _, cartItem := range cartResult.Items {
		productResp, resultErr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{
			Id: cartItem.ProductId,
		})

		if resultErr != nil {
			return nil, kerrors.NewGRPCBizStatusError(5005002, resultErr.Error())
		}

		if productResp.Product == nil {
			continue
		}

		p := productResp.Product.Price

		cost := p * float32(cartItem.Quantity)
		total += cost

		items = append(items, &order.OrderItem{
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
			Street:  req.Address.Street,
			City:    req.Address.City,
			State:   req.Address.State,
			Country: req.Address.Country,
			Zip:     req.Address.Zip,
		},
		Items: items,
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005003, err.Error())
	}

	if orderResp != nil && orderResp.Order != nil {
		orderId = orderResp.Order.OrderId
	}

	// u, _ := uuid.NewRandom()

	// orderId = u.String()

	payReq := &payment.ChargeReq{
		OrderId: orderId,
		UserId:  req.UserId,
		Amount:  total,
		Card: &payment.CreditCard{
			Id:              req.CreditCard.Id,
			Cvv:             req.CreditCard.Cvv,
			ExpirationYear:  req.CreditCard.ExpirationYear,
			ExpirationMonth: req.CreditCard.ExpirationMonth,
		},
	}

	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005003, err.Error())
	}

	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005004, err.Error())
	}

	data, _ := proto.Marshal(&email.EmailReq{
		From:        "checkout@example.com",
		To:          req.Email,
		ContentType: "text/plain",
		Subject:     "Order Confirmation",
		Content:     "Order Confirmation",
	})

	msg := &nats.Msg{
		Subject: "email",
		Data:    data,
		Header:  make(nats.Header),
	}
	otel.GetTextMapPropagator().Inject(s.ctx, propagation.HeaderCarrier(msg.Header))

	_ = mq.Nc.PublishMsg(msg)

	klog.Info(paymentResult)
	klog.Info(orderResp)

	resp = &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResult.TransactionId,
	}

	return
}
