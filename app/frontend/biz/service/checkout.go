package service

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	common "github.com/spark4862/smartmall/app/frontend/hertz_gen/frontend/common"
	"github.com/spark4862/smartmall/app/frontend/infra/rpc"
	frontendUtils "github.com/spark4862/smartmall/app/frontend/utils"
	rpcCart "github.com/spark4862/smartmall/rpc_gen/kitex_gen/cart"
	rpcProduct "github.com/spark4862/smartmall/rpc_gen/kitex_gen/product"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	var items []map[string]string
	userId := frontendUtils.GetUserIdFromCtx(h.Context)

	carts, err := rpc.CartClient.GetCart(h.Context, &rpcCart.GetCartReq{
		UserId: uint32(userId),
	})
	if err != nil {
		return
	}

	var total float32

	for _, v := range carts.Items {
		// 给前端渲染的
		// 但是感觉逻辑有重复，前端渲染获取了一遍数据，后端算价格又获取了一遍

		productResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcProduct.GetProductReq{
			Id: v.ProductId,
		})
		if err != nil {
			return nil, err
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product
		items = append(items, map[string]string{
			"Name":        p.Name,
			"Description": p.Description,
			"Price":       strconv.FormatFloat(float64(p.Price), 'f', 2, 64),
			"Qty":         strconv.Itoa(int(v.Quantity)),
		})
		total += float32(v.Quantity) * p.Price
	}

	return utils.H{
		"title": "Checkout",
		"items": items,
		"total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
