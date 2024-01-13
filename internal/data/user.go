package data

import (
	"context"

	"kratos-realworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *UserRepo) Register(ctx context.Context, register *biz.UserRegister) (*biz.User, error) {
	return nil, ErrNotImplemented
}

func (r *UserRepo) Login(ctx context.Context, login *biz.UserLogin) (*biz.User, error) {
	return nil, ErrNotImplemented
}

func (r *UserRepo) Update(ctx context.Context, update *biz.UserUpdate) (*biz.User, error) {
	return nil, ErrNotImplemented
}

func (r *UserRepo) GetCurrentUser(ctx context.Context) (*biz.User, error) {
	return nil, ErrNotImplemented
}
