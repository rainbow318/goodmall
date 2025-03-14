package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	frontendUtils "github.com/suutest/app/frontend/utils"
)

// GlobalAuth中间件在每个请求中从session获取用户ID并存入context
func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) { // 便于业务逻辑获取用户身份相关的内容
		// 从session中获取用户信息，然后放在context里
		s := sessions.Default(c)
		ctx = context.WithValue(ctx, frontendUtils.SessionUserId, s.Get("user_id")) // 这样，其他业务逻辑需要user_id时就可以直接从context里取，而不是还要用session
		c.Next(ctx)
	}
}

// 利用Hertz的中间件写鉴权的逻辑
// Auth中间件检查session中是否存在user_id，如果不存在则重定向到登录页面
func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		s := sessions.Default(c)   // 从cookie中提取sessionid
		userId := s.Get("user_id") // 通过sessionid从redis中获取数据

		// 如果session中没有user_id，就表明未登录，所以重定向到登录页面
		if userId == nil {
			c.Redirect(302, []byte("/sign-in?next="+c.FullPath()))
			c.Abort()
			return
		}
		// 验证通过
		c.Next(ctx)
	}
}
