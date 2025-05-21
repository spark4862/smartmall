package service

import (
	"context"

	"github.com/spark4862/smartmall/app/product/biz/dal/mysql"
	"github.com/spark4862/smartmall/app/product/biz/model"
	product "github.com/spark4862/smartmall/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// Finish your business logic.

	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	products, err := productQuery.SearchProducts(req.Query)
	var results []*product.Product
	for _, p := range products {
		results = append(results, &product.Product{
			Id:          int32(p.ID),
			Name:        p.Name,
			Description: p.Description,
			Picture:     p.Picture,
			Price:       p.Price,
		})
	}

	return &product.SearchProductsResp{
		Results: results,
	}, err
}
