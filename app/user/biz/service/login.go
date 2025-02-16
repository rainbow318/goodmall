package service

import (
	"context"
	"errors"

	"github.com/suutest/app/user/biz/dal/mysql"
	"github.com/suutest/app/user/biz/model"
	user "github.com/suutest/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	// 业务逻辑：查找账号密码
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email of password is empty")
	}
	row, err := model.GetByEmail(s.ctx, mysql.DB, req.Email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(row.PasswordHashed), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	resp = &user.LoginResp{
		UserId: int32(row.ID),
	}
	return resp, nil
}
