package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/spark4862/smartmall/app/cart/biz/dal/mysql"
	"github.com/spark4862/smartmall/app/cart/model"
	"github.com/spark4862/smartmall/app/cart/rpc"
	cart "github.com/spark4862/smartmall/rpc_gen/kitex_gen/cart"
	"github.com/spark4862/smartmall/rpc_gen/kitex_gen/product"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// Finish your business logic.
	productResp, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{
		Id: req.Item.ProductId,
	})
	if err != nil {
		return
	}
	if productResp.Product == nil || productResp.Product.Id == 0 {
		return nil, kerrors.NewBizStatusError(5001001, "product not found")
	}

	item := model.Cart{
		ProductId: req.Item.ProductId,
		Qty:       req.Item.Quantity,
		UserId:    req.UserId,
	}

	err = model.AddItem(s.ctx, mysql.DB, &item)
	if err != nil {
		return nil, kerrors.NewBizStatusError(5001002, "product not found")
	}

	return &cart.AddItemResp{}, nil
}
