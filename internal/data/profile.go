package data

import (
	"context"

	"kratos-realworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type Following struct {
	gorm.Model
	UserID      uint
	User        User
	FollowingID uint
	Following   User
}

type ProfileRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo
func NewProfileRepo(data *Data, logger log.Logger) biz.ProfileRepo {
	return &ProfileRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *ProfileRepo) Follow(ctx context.Context, userID, followID uint) error {
	return ErrNotImplemented
}
func (r *ProfileRepo) UnFollow(ctx context.Context, userID, followID uint) error {
	return ErrNotImplemented
}
func (r *ProfileRepo) GetFollowers(ctx context.Context, userID uint) ([]*biz.Profile, error) {
	return nil, ErrNotImplemented
}
func (r *ProfileRepo) GetFollowings(ctx context.Context, userID uint) ([]*biz.Profile, error) {
	return nil, ErrNotImplemented
}
