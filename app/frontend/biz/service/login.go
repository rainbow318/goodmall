package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	auth "github.com/suutest/app/frontend/hertz_gen/frontend/auth"
	"github.com/suutest/app/frontend/infra/rpc"
	"github.com/suutest/rpc_gen/kitex_gen/user"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (redirect string, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// TODO 用户提交过来后进行一个验证
	resp, err := rpc.UserClient.Login(h.Context, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return "", err
	}
	// 获取与当前请求关联的session（由中间件自动管理）
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", resp.UserId) // 在已存在的session中存储用户凭证
	err = session.Save()                // 将修改持久化到redis
	if err != nil {
		return "", err
	}
	redirect = "/"
	if req.Next != "" {
		redirect = req.Next
	}
	return redirect, nil
}
