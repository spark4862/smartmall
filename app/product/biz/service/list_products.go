package service

import (
	"context"

	"github.com/spark4862/smartmall/app/product/biz/dal/mysql"
	"github.com/spark4862/smartmall/app/product/biz/model"
	"github.com/spark4862/smartmall/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductReq) (resp *product.ListProductResp, err error) {
	// Finish your business logic.
	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)

	categories, err := categoryQuery.GetProductsByCategoryName(req.CategoryName)
	resp = &product.ListProductResp{}
	for _, c := range categories {
		for _, p := range c.Products {
			resp.Products = append(resp.Products, &product.Product{
				Id:          int32(p.ID),
				Name:        p.Name,
				Description: p.Description,
				Picture:     p.Picture,
				Price:       p.Price,
			})
		}
	}
	return
}
