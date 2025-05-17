package service

import (
	"context"
	product "github.com/spark4862/smartmall/rpc_gen/kitex_gen/product"
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

	return
}
