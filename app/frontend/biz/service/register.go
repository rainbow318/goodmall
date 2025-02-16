package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	auth "github.com/suutest/app/frontend/hertz_gen/frontend/auth"
	common "github.com/suutest/app/frontend/hertz_gen/frontend/common"
	"github.com/suutest/app/frontend/infra/rpc"
	"github.com/suutest/rpc_gen/kitex_gen/user"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()

	// 调用user服务的Register方法
	userResp, err := rpc.UserClient.Register(h.Context, &user.RegisterReq{
		Email:           req.Email,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
	})
	if err != nil {
		return nil, err
	}
	session := sessions.Default(h.RequestContext) // 从请求上下文中获取一个会话对象
	session.Set("user_id", userResp.UserId)       // 将user服务返回中的user_id保存在会话中
	err = session.Save()                          // 将会话信息保存到存储中（Redis）
	if err != nil {
		return nil, err
	}
	return
}
