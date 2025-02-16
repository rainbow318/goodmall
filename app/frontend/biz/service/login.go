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

	session := sessions.Default(h.RequestContext)
	session.Set("user_id", resp.UserId) // 将user_id保存到当前会话中。此时并没有根据传入的 req 来进行验证或查找数据库，直接将user_id设置为 1。
	err = session.Save()                // 将上面设置的 user_id 存储到会话中，这样在随后的请求中，服务器可以通过会话中的 user_id 来识别用户的身份
	if err != nil {
		return "", err
	}
	redirect = "/"
	if req.Next != "" {
		redirect = req.Next
	}
	return redirect, nil
}
