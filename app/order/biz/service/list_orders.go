package service

import (
	"context"

	"github.com/spark4862/smartmall/app/order/biz/dal/mysql"
	"github.com/spark4862/smartmall/app/order/biz/model"
	cartRpc "github.com/spark4862/smartmall/rpc_gen/kitex_gen/cart"
	orderRpc "github.com/spark4862/smartmall/rpc_gen/kitex_gen/order"
)

type ListOrdersService struct {
	ctx context.Context
} // NewListOrdersService new ListOrdersService
func NewListOrdersService(ctx context.Context) *ListOrdersService {
	return &ListOrdersService{ctx: ctx}
}

// Run create note info
func (s *ListOrdersService) Run(req *orderRpc.ListOrdersReq) (resp *orderRpc.ListOrdersResp, err error) {
	// Finish your business logic.

	list, err := model.ListOrder(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return
	}

	var orders []*orderRpc.Order
	for _, v := range list {
		var items []*orderRpc.OrderItem
		for _, o := range v.OrderItems {
			items = append(items, &orderRpc.OrderItem{
				Item: &cartRpc.CartItem{
					ProductId: o.ProductId,
					Quantity:  o.Quantity,
				},
				Cost: o.Cost,
			})
		}

		orders = append(orders, &orderRpc.Order{
			OrderId:  v.OrderId,
			UserId:   v.UserId,
			Currency: v.Currency,
			Address: &orderRpc.Address{
				City:    v.Consignee.City,
				Country: v.Consignee.Country,
				State:   v.Consignee.State,
				Street:  v.Consignee.Street,
				Zip:     v.Consignee.ZipCode,
			},
			Email: v.Consignee.Email,
			Items: items,
		})
	}
	resp = &orderRpc.ListOrdersResp{
		Orders: orders,
	}

	return
}
