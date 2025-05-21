package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	category "github.com/spark4862/smartmall/app/frontend/hertz_gen/frontend/category"
	"github.com/spark4862/smartmall/app/frontend/infra/rpc"
	rpcproduct "github.com/spark4862/smartmall/rpc_gen/kitex_gen/product"
)

type GetCategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCategoryService(Context context.Context, RequestContext *app.RequestContext) *GetCategoryService {
	return &GetCategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCategoryService) Run(req *category.CategoryReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	listProductResp, err := rpc.ProductClient.ListProducts(h.Context, &rpcproduct.ListProductReq{CategoryName: req.Category})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"title": "Category",
		"items": listProductResp.Products,
	}, nil
}
