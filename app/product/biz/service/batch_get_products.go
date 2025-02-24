package service

import (
	"context"

	"github.com/suutest/app/product/biz/dal/mysql"
	"github.com/suutest/app/product/biz/dal/redis"
	"github.com/suutest/app/product/biz/model"
	product "github.com/suutest/rpc_gen/kitex_gen/product"
)

type BatchGetProductsService struct {
	ctx context.Context
} // NewBatchGetProductsService new BatchGetProductsService
func NewBatchGetProductsService(ctx context.Context) *BatchGetProductsService {
	return &BatchGetProductsService{ctx: ctx}
}

// Run create note info
func (s *BatchGetProductsService) Run(req *product.BatchGetProductsReq) (resp *product.BatchGetProductsResp, err error) {
	// productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	// ps, err := productQuery.BatchGetByIds(req.Ids)
	productQuery := model.NewCachedProductQuery(s.ctx, mysql.DB, redis.RedisClient)
	ps, err := productQuery.BatchGetByIds(req.Ids)
	if err != nil {
		return nil, err
	}
	products := make(map[uint32]*product.Product)
	for _, p := range ps {
		products[uint32(p.ID)] = &product.Product{
			Id:          uint32(p.ID),
			Name:        p.Name,
			Description: p.Description,
			Picture:     p.Picture,
			Price:       p.Price,
		}
	}
	return &product.BatchGetProductsResp{
		Products: products,
	}, nil
}
