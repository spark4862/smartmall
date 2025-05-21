package service

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	common "github.com/spark4862/smartmall/app/frontend/hertz_gen/frontend/common"
	"github.com/spark4862/smartmall/app/frontend/infra/rpc"
	frontendUtils "github.com/spark4862/smartmall/app/frontend/utils"
	"github.com/spark4862/smartmall/rpc_gen/kitex_gen/cart"
	"github.com/spark4862/smartmall/rpc_gen/kitex_gen/product"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	getCartResp, err := rpc.CartClient.GetCart(h.Context, &cart.GetCartReq{
		UserId: uint32(frontendUtils.GetUserIdFromCtx(h.Context)),
	})
	if err != nil {
		return
	}
	var items []map[string]string
	var total float64
	for _, item := range getCartResp.Items {
		getProductResp, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductReq{
			Id: item.ProductId,
		})
		if err != nil {
			continue
		}
		p := getProductResp.Product
		items = append(items, map[string]string{
			"Name":        p.Name,
			"Description": p.Description,
			"Price":       strconv.FormatFloat(float64(p.Price), 'f', 2, 64),
			"Picture":     p.Picture,
			"Qty":         strconv.Itoa(int(item.Quantity)),
		})
		total += float64(p.Price) * float64(item.Quantity)
	}
	return utils.H{
		"title": "Cart",
		"items": items,
		"total": strconv.FormatFloat(total, 'f', 2, 64),
	}, nil
}
