package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/spark4862/smartmall/app/product/biz/dal/mysql"
	"github.com/spark4862/smartmall/app/product/biz/model"
	product "github.com/spark4862/smartmall/rpc_gen/kitex_gen/product"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// Finish your business logic.
	if req.Id == 0 {
		return nil, kerrors.NewGRPCBizStatusError(2004001, "product id is required")
	}

	productQuery := model.NewProductQuery(s.ctx, mysql.DB)

	p, err := productQuery.GetByID(int(req.Id))
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(2004002, "product not found")
	}

	return &product.GetProductResp{
		Product: &product.Product{
			Id:          int32(p.ID),
			Name:        p.Name,
			Description: p.Description,
			Picture:     p.Picture,
			Price:       p.Price,
		},
	}, nil
}
