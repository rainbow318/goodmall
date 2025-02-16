package middleware

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

// 注册一个中间件
func Register(h *server.Hertz) {
	h.Use(GlobalAuth())
}
