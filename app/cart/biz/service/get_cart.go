package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/spark4862/smartmall/app/cart/biz/dal/mysql"
	"github.com/spark4862/smartmall/app/cart/model"
	cart "github.com/spark4862/smartmall/rpc_gen/kitex_gen/cart"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// Finish your business logic.
	items, err := model.GetCartByUserId(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(5003001, err.Error())
	}
	resp = &cart.GetCartResp{}

	for _, item := range items {
		resp.Items = append(resp.Items, &cart.CartItem{
			ProductId: item.ProductId,
			Quantity:  item.Qty,
		})
	}
	return
}
