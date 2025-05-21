package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/google/uuid"
	"github.com/spark4862/smartmall/app/order/biz/dal/mysql"
	"github.com/spark4862/smartmall/app/order/biz/model"
	order "github.com/spark4862/smartmall/rpc_gen/kitex_gen/order"
	"gorm.io/gorm"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	// Finish your business logic.
	if len(req.Items) == 0 {
		err = kerrors.NewBizStatusError(6004001, "item is empty")
		return
	}

	err = mysql.DB.Transaction(func(db *gorm.DB) error {
		orderId, _ := uuid.NewRandom()

		o := &model.Order{
			OrderId:  orderId.String(),
			UserId:   req.UserId,
			Currency: req.Currency,
			Consignee: model.Consignee{
				Email: req.Email,
			},
		}
		if req.Address != nil {
			a := req.Address
			o.Consignee.Street = a.Street
			o.Consignee.City = a.City
			o.Consignee.State = a.State
			o.Consignee.Country = a.Country
			o.Consignee.ZipCode = a.Zip
		}

		if err := db.Create(o).Error; err != nil {
			return err
		}
		var items []model.OrderItem
		for _, v := range req.Items {
			items = append(items, model.OrderItem{
				OrderId:   o.OrderId,
				ProductId: int32(v.Item.ProductId),
				Quantity:  int32(v.Item.Quantity),
				Cost:      v.Cost,
			})
		}
		if err := db.Create(&items).Error; err != nil {
			return err
		}

		resp = &order.PlaceOrderResp{
			Order: &order.OrderResult{
				OrderId: o.OrderId,
			},
		}

		return nil
	})

	return
}
