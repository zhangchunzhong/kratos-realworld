package biz

import (
	"context"
	v1 "kratos-realworld/api/conduit/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type User struct {
	Id       uint
	Email    string
	Username string
	PwdHash  string
	Bio      string
	Image    string
}

type UserLogin struct {
	Email    string
	Token    string
	Password string
}

type UserRegister struct {
	Email    string
	Username string
	Password string
}

type UserUpdate struct {
	Email    string
	Username string
	Password string
	Bio      string
	Image    string
}

type Author Profile

// UserRepo is a User repo.
type UserRepo interface {
	Register(ctx context.Context, register *UserRegister) (*User, error)
	Login(ctx context.Context, login *UserLogin) (*User, error)
	Update(ctx context.Context, update *UserUpdate) (*User, error)
	GetCurrentUser(ctx context.Context) (*User, error)
}
