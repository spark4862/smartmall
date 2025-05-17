package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/spark4862/smartmall/app/frontend/hertz_gen/frontend/common"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	resp := make(map[string]any)

	items := []map[string]any{
		{"Name": "egg", "Price": 10, "Picture": "/static/img/food/egg.jpg"},
		{"Name": "oatmeal", "Price": 10, "Picture": "/static/img/food/oatmeal.jpg"},
		{"Name": "shrimp", "Price": 10, "Picture": "/static/img/food/shrimp.jpg"},
		{"Name": "pomfret", "Price": 10, "Picture": "/static/img/food/pomfret.jpg"},
		{"Name": "shrimp", "Price": 10, "Picture": "/static/img/food/shrimp.jpg"},
		{"Name": "egg", "Price": 10, "Picture": "/static/img/food/egg.jpg"},
	}

	resp["title"] = "Smart Mall"
	resp["items"] = items

	return resp, nil
}
