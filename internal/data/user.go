package data

import (
	"context"

	"kratos-realworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string `gorm:"size:500"`
	Username     string `gorm:"size:500"`
	Bio          string `gorm:"size:1000"`
	Image        string `gorm:"size:1000"`
	PasswordHash string `gorm:"size:500"`
	Following    uint32
}

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

func (r *UserRepo) GetCurrent(ctx context.Context) (*biz.User, error) {
	return nil, ErrNotImplemented
}

func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	return nil, ErrNotImplemented
}

func (r *UserRepo) GetUserByUsername(ctx context.Context, username string) (*biz.User, error) {
	return nil, ErrNotImplemented
}

func (r *UserRepo) GetUserById(ctx context.Context, id uint) (*biz.User, error) {
	return nil, ErrNotImplemented
}
