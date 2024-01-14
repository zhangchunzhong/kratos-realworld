package biz

import (
	"context"
	v1 "kratos-realworld/api/conduit/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

var (
	// ErrProfileNotFound is profile not found.
	ErrProfileNotFound = errors.NotFound(v1.ErrorReason_PROFILE_NOT_FOUND.String(), "profile not found")
)

type Profile struct {
	Username  string
	Bio       string
	Image     string
	Following bool
}

type ProfileRepo interface {
	Follow(ctx context.Context, userID, followID uint) error
	UnFollow(ctx context.Context, userID, followID uint) error
	GetFollowers(ctx context.Context, userID uint) ([]*Profile, error)
	GetFollowings(ctx context.Context, userID uint) ([]*Profile, error)
}

// Follow a user
//
//	(1) Authentication optional
func (uc *ConduitUseCase) Follow(ctx context.Context, userID, followID uint) error {
	return uc.pr.Follow(ctx, userID, followID)
}

// UnFollow a user
//
//	(1) Authentication optional
func (uc *ConduitUseCase) UnFollow(ctx context.Context, userID, followID uint) error {
	return uc.pr.UnFollow(ctx, userID, followID)
}

// Get followers of a user
//
//	(1) Authentication optional, returns a Profile[]
func (uc *ConduitUseCase) GetFollowers(ctx context.Context, userID uint) ([]*Profile, error) {
	return uc.pr.GetFollowers(ctx, userID)
}

// Get followings of a user
//
//	(1) Authentication optional, returns a Profile[]
func (uc *ConduitUseCase) GetFollowings(ctx context.Context, userID uint) ([]*Profile, error) {
	return uc.pr.GetFollowings(ctx, userID)
}
