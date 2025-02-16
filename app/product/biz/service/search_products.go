package service

import (
	"context"

	"github.com/suutest/app/product/biz/dal/mysql"
	"github.com/suutest/app/product/biz/model"
	product "github.com/suutest/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductReq) (resp *product.SearchProductResp, err error) {
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	products, err := productQuery.SearchProduct(req.Query)
	var results []*product.Product
	// resp = &product.SearchProductResp{}
	for _, v := range products {
		//resp.Results = append(resp.Results, &product.Product{
		//	Id:          v.ID,
		//	Name:        v.Name,
		//	Description: v.Description,
		//	Picture:     v.Picture,
		//	Price:       v.Price,
		//})
		results = append(results, &product.Product{
			Id:          uint32(v.ID),
			Name:        v.Name,
			Description: v.Description,
			Picture:     v.Picture,
			Price:       v.Price,
		})
	}
	return &product.SearchProductResp{Results: results}, nil
}
