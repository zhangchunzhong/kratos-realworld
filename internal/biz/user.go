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
	ID       uint
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
	GetCurrent(ctx context.Context) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	GetUserById(ctx context.Context, id uint) (*User, error)
}

// Authentication Header:
//
//	(1) You can read the authentication header from the headers of the request
//	(2)Authorization: Token jwt.token.here

// Registration
//
// (1) No authentication required, returns a User
// (2) Required fields: email, username, password
func (uc *ConduitUseCase) UserRegister(ctx context.Context, register *UserRegister) (*User, error) {
	return uc.ur.Register(ctx, register)
}

// Login
//
// (1) No authentication required, returns a User
// (2) Required fields: email, password
func (uc *ConduitUseCase) UserLogin(ctx context.Context, login *UserLogin) (*User, error) {
	return uc.ur.Login(ctx, login)
}

// Update User
//
// (1) Authentication required, returns the User
// (2) Optional fields: email, username, password, image, bio
func (uc *ConduitUseCase) UserUpdate(ctx context.Context, update *UserUpdate) (*User, error) {
	return uc.ur.Update(ctx, update)
}

// Get Current User
//
// (1) Authentication required, returns a User that's the current user
func (uc *ConduitUseCase) UserGetCurrent(ctx context.Context) (*User, error) {
	return uc.ur.GetCurrent(ctx)
}

// Get User by Email
//
// (1) Authentication optional, returns a User
func (uc *ConduitUseCase) UserGetByEmail(ctx context.Context, email string) (*User, error) {
	return uc.ur.GetUserByEmail(ctx, email)
}

// Get User by Username
//
// (1) Authentication optional, returns a User
func (uc *ConduitUseCase) UserGetByUsername(ctx context.Context, username string) (*User, error) {
	return uc.ur.GetUserByUsername(ctx, username)
}

// Get User by Id
//
// (1) Authentication optional, returns a User
func (uc *ConduitUseCase) UserGetById(ctx context.Context, id uint) (*User, error) {
	return uc.ur.GetUserById(ctx, id)
}
