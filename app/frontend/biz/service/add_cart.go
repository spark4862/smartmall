package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/spark4862/smartmall/app/frontend/hertz_gen/frontend/cart"
	common "github.com/spark4862/smartmall/app/frontend/hertz_gen/frontend/common"
	"github.com/spark4862/smartmall/app/frontend/infra/rpc"
	frontendUtils "github.com/spark4862/smartmall/app/frontend/utils"
	rpccart "github.com/spark4862/smartmall/rpc_gen/kitex_gen/cart"
)

type AddCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartService(Context context.Context, RequestContext *app.RequestContext) *AddCartService {
	return &AddCartService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartService) Run(req *cart.AddCartReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	_, err = rpc.CartClient.AddItem(h.Context, &rpccart.AddItemReq{
		UserId: uint32(frontendUtils.GetUserIdFromCtx(h.Context)),
		Item: &rpccart.CartItem{
			ProductId: int32(req.ProductId),
			Quantity:  req.ProductNum,
		},
	})
	if err != nil {
		return
	}

	return
}
