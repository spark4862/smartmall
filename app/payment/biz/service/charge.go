package service

import (
	"context"
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/kerrors"
	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
	"github.com/spark4862/smartmall/app/payment/biz/dal/mysql"
	"github.com/spark4862/smartmall/app/payment/biz/model"
	payment "github.com/spark4862/smartmall/rpc_gen/kitex_gen/payment"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	// Finish your business logic.
	card := creditcard.Card{
		Number: req.Card.Id,
		Cvv:    strconv.Itoa(int(req.Card.Cvv)),
		Month:  strconv.Itoa(int(req.Card.ExpirationMonth)),
		Year:   strconv.Itoa(int(req.Card.ExpirationYear)),
	}

	err = card.Validate(true)
	if err != nil {
		// 400 服务代号 4001 参数错误
		// 400 客户端 500 服务端
		return nil, kerrors.NewGRPCBizStatusError(4004001, err.Error())
	}

	transactionId, err := uuid.NewRandom()
	if err != nil {
		return nil, kerrors.NewBizStatusError(4005001, err.Error())
	}

	err = model.CreatePaymentLog(s.ctx, mysql.DB, &model.PaymentLog{
		UserId:        req.UserId,
		OrderId:       req.OrderId,
		TransactionId: transactionId.String(),
		Amount:        req.Amount,
		PayAt:         time.Now(),
	})
	if err != nil {
		return nil, kerrors.NewBizStatusError(4005002, err.Error())
	}

	return &payment.ChargeResp{
		TransactionId: transactionId.String(),
	}, nil
}
