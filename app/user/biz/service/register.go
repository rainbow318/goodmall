package service

import (
	"context"
	"errors"

	"github.com/suutest/app/user/biz/dal/mysql"
	"github.com/suutest/app/user/biz/model"
	"github.com/suutest/app/user/infra/filter"
	user "github.com/suutest/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.
	if req.Email == "" || req.Password == "" || req.PasswordConfirm == "" {
		return nil, errors.New("email or password is empty")
	}
	if filter.FilterState && filter.Filter.TestString(req.Email) {
		return nil, errors.New("user is exist")
	}
	if req.Password != req.PasswordConfirm {
		return nil, errors.New("password not match")
	}
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	// 严格来说的话，这里的email也要校验一下
	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(passwordHashed),
	}
	err = model.Create(s.ctx, mysql.DB, newUser)
	if err != nil {
		return nil, err
	}
	filter.FilterMutex.Lock()
	defer filter.FilterMutex.Unlock()
	filter.Filter.AddString(req.Email)
	return &user.RegisterResp{UserId: int32(newUser.ID)}, nil
}
