package service

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	common "github.com/spark4862/smartmall/app/frontend/hertz_gen/frontend/common"
	"github.com/spark4862/smartmall/app/frontend/infra/rpc"
	"github.com/spark4862/smartmall/app/frontend/types"
	frontendUtils "github.com/spark4862/smartmall/app/frontend/utils"
	orderRpc "github.com/spark4862/smartmall/rpc_gen/kitex_gen/order"
	productRpc "github.com/spark4862/smartmall/rpc_gen/kitex_gen/product"
)

type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	orderResp, err := rpc.OrderClient.ListOrders(h.Context, &orderRpc.ListOrdersReq{UserId: uint32(userId)})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var list []types.Order

	for _, order := range orderResp.Orders {
		var (
			total float32
			items []types.OrderItem
		)
		for _, item := range order.Items {
			total += item.Cost

			productResp, err := rpc.ProductClient.GetProduct(h.Context, &productRpc.GetProductReq{Id: item.Item.ProductId})
			if err != nil {
				continue
			}
			if productResp == nil || productResp.Product == nil {
				continue
			}
			p := productResp.Product
			items = append(items, types.OrderItem{
				ProductName: p.Name,
				Qty:         uint32(item.Item.Quantity),
				Cost:        item.Cost,
				Picture:     p.Picture,
			})
		}
		created := time.Unix(int64(order.CreatedAt), 0)
		list = append(list, types.Order{
			OrderId:     order.OrderId,
			Cost:        total,
			Items:       items,
			CreatedDate: created.Format("2006-01-02 15:04:05"),
		})
	}

	return utils.H{
		"title":  "Order",
		"orders": list,
	}, nil
}
