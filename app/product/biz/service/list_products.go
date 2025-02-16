package service

import (
	"context"

	"github.com/suutest/app/product/biz/dal/mysql"
	"github.com/suutest/app/product/biz/model"
	product "github.com/suutest/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// 通过一个分类名称获取到一个商品列表
	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)
	c, err := categoryQuery.GetProductsByCategoryName(req.CategoriyName)
	resp = &product.ListProductsResp{}
	for _, v1 := range c { // 遍历检索到的类别
		for _, v := range v1.Products { // 遍历该类别下的所有商品
			resp.Products = append(resp.Products, &product.Product{
				Id:          uint32(v.ID),
				Name:        v.Name,
				Description: v.Description,
				Picture:     v.Picture,
				Price:       v.Price,
			})
		}
	}
	return resp, nil
}
