package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	frontendUtils "github.com/suutest/app/frontend/utils"
)

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) { // 便于业务逻辑获取用户身份相关的内容
		// 从session中获取用户信息，然后放在context里
		s := sessions.Default(c)
		ctx = context.WithValue(ctx, frontendUtils.SessionUserId, s.Get("user_id")) // 这样，其他业务逻辑需要user_id时就可以直接从context里取，而不是还要用session
		c.Next(ctx)
	}
}

// 利用Hertz的中间件写鉴权的逻辑
func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		s := sessions.Default(c)
		userId := s.Get("user_id") // 从cookie中尝试获取用户数据，如果获取不到，就让他跳转

		if userId == nil {
			c.Redirect(302, []byte("/sign-in?next="+c.FullPath()))
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}
