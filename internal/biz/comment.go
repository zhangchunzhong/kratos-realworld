package biz

import (
	"context"
	v1 "kratos-realworld/api/conduit/v1"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
)

var (
	// ErrCommentNotFound is comment not found.
	ErrCommentNotFound = errors.NotFound(v1.ErrorReason_COMMENT_NOT_FOUND.String(), "comment not found")
)

type Comment struct {
	ID        uint
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time

	ArticleId uint
	Article   *Article
}

type CommentRepo interface {
	Add(ctx context.Context, slug string, comment *Comment) (*Comment, error)
	Delete(ctx context.Context, slug string, CommentId string) error
	Get(ctx context.Context, slug string) ([]*Comment, error)
}

// Add Comment
//
//	(1) Authentication required, returns the created Comment
//	(2) Required fields: body
func (uc *ConduitUseCase) AddComment(ctx context.Context, slug string, comment *Comment) (*Comment, error) {
	return uc.cr.Add(ctx, slug, comment)
}

// Delete Comment
//
// (1) Authentication required
// (2) Returns an empty response
func (uc *ConduitUseCase) DeleteComment(ctx context.Context, slug string, CommentId string) error {
	return uc.cr.Delete(ctx, slug, CommentId)
}

// Get Comments from an Article
//
//	(1) Authentication optional, returns multiple comments
func (uc *ConduitUseCase) GetComment(ctx context.Context, slug string) ([]*Comment, error) {
	return uc.cr.Get(ctx, slug)
}
