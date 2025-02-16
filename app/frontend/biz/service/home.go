package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/suutest/app/frontend/hertz_gen/frontend/common"
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
	resp := make(map[string]any)
	item := []map[string]any{
		{"Name": "Macbook", "Price": 7999, "Picture": "/static/image/laptop-1.avif"},
		{"Name": "Airpods", "Price": 750, "Picture": "/static/image/headphone-1.avif"},
		{"Name": "Airpods Max", "Price": 3500, "Picture": "/static/image/headphone-2.avif"},
		{"Name": "iPad", "Price": 5000, "Picture": "/static/image/tablet-1.avif"},
		{"Name": "Watch", "Price": 4500, "Picture": "/static/image/watch-1.avif"},
		{"Name": "Watch", "Price": 2500, "Picture": "/static/image/watch-2.avif"},
		{"Name": "iPhone 15", "Price": 5999, "Picture": "/static/image/iphone-1.avif"},
		{"Name": "iPhone 15", "Price": 5999, "Picture": "/static/image/iphone-2.avif"},
	}
	resp["Title"] = "Hot Sales"
	resp["Items"] = item
	return resp, nil
}
