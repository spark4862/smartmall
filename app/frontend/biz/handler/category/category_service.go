package category

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/spark4862/smartmall/app/frontend/biz/service"
	"github.com/spark4862/smartmall/app/frontend/biz/utils"
	category "github.com/spark4862/smartmall/app/frontend/hertz_gen/frontend/category"
)

// GetCategory .
// @router /category/:category [GET]
func GetCategory(ctx context.Context, c *app.RequestContext) {
	var err error
	var req category.CategoryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	resp, err := service.NewGetCategoryService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "category", utils.WarpResponse(ctx, c, resp))
}
