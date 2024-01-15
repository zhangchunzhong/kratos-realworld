package data

import (
	"context"

	"kratos-realworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ArticleSlug string
	Article     Article `gorm:"references:Slug"`
	Body        string
	AuthorID    uint
}

type CommentRepo struct {
	data *Data
	log  *log.Helper
}

// NewCommentRepo
func NewCommentRepo(data *Data, logger log.Logger) biz.CommentRepo {
	return &CommentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *CommentRepo) Add(ctx context.Context, slug string, comment *biz.Comment) (*biz.Comment, error) {
	return nil, ErrNotImplemented
}
func (r *CommentRepo) Delete(ctx context.Context, slug string, CommentId string) error {
	return ErrNotImplemented
}
func (r *CommentRepo) Get(ctx context.Context, slug string) ([]*biz.Comment, error) {
	return nil, ErrNotImplemented
}
