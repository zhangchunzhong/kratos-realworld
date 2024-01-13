package data

import (
	"context"
	"kratos-realworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type ArticleRepo struct {
	data *Data
	log  *log.Helper
}

// NewArticleRepo
func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &ArticleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *ArticleRepo) Create(ctx context.Context, article *biz.Article) (*biz.Article, error) {
	return nil, ErrNotImplemented
}
func (r *ArticleRepo) Get(ctx context.Context, slug string) (*biz.Article, error) {
	return nil, ErrNotImplemented
}
func (r *ArticleRepo) Update(ctx context.Context, slug string, article *biz.Article) (*biz.Article, error) {
	return nil, ErrNotImplemented
}
func (r *ArticleRepo) Delete(ctx context.Context, slug string) error {
	return ErrNotImplemented
}
func (r *ArticleRepo) List(ctx context.Context, opt *biz.ListOptions) ([]*biz.Article, error) {
	return nil, ErrNotImplemented
}
func (r *ArticleRepo) Feed(ctx context.Context, opt *biz.ListOptions, userId uint) ([]*biz.Article, error) {
	return nil, ErrNotImplemented
}
func (r *ArticleRepo) Favorite(ctx context.Context, slug string, userID uint) error {
	return ErrNotImplemented
}
func (r *ArticleRepo) UnFavorite(ctx context.Context, slug string, userID uint) error {
	return ErrNotImplemented
}
func (r *ArticleRepo) GetTags(ctx context.Context) ([]*biz.Tag, error) {
	return nil, ErrNotImplemented
}
