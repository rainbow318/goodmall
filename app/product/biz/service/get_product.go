package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/suutest/app/product/biz/dal/mysql"
	"github.com/suutest/app/product/biz/dal/redis"
	"github.com/suutest/app/product/biz/model"
	product "github.com/suutest/rpc_gen/kitex_gen/product"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	if req.Id == 0 {
		return nil, kerrors.NewGRPCBizStatusError(2004001, "produce id is required")
	}
	// productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	productQuery := model.NewCachedProductQuery(s.ctx, mysql.DB, redis.RedisClient)
	// p, err := productQuery.GetById(int(req.Id))
	p, err := productQuery.TwoLevelCacheGetById(int(req.Id))
	if err != nil {
		return nil, err
	}
	return &product.GetProductResp{
		Product: &product.Product{
			Id:          uint32(p.ID),
			Name:        p.Name,
			Description: p.Description,
			Picture:     p.Picture,
			Price:       p.Price,
		},
	}, nil
}
