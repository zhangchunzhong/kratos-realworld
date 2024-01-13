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
	Id        uint
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
