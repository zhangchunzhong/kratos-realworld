package biz

import (
	"context"
	v1 "kratos-realworld/api/conduit/v1"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
)

var (
	// ErrArticleNotFound is article not found.
	ErrArticleNotFound = errors.NotFound(v1.ErrorReason_ARTICLE_NOT_FOUND.String(), "article not found")
)

type ListOptions struct {
	Author    string
	Favorited string
	Tag       string
	Offset    int64
	Limit     int64
	UserID    uint
}

type Article struct {
	Id             uint
	Slug           string
	Title          string
	Description    string
	Body           string
	TagList        []string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Favorited      bool
	FavoritesCount uint

	AuthorId uint
	Author   *Author
}

type Tag string

type ArticleRepo interface {
	Create(ctx context.Context, article *Article) (*Article, error)
	Get(ctx context.Context, slug string) (*Article, error)
	Update(ctx context.Context, slug string, article *Article) (*Article, error)
	Delete(ctx context.Context, slug string) error
	List(ctx context.Context, opt *ListOptions) ([]*Article, error)
	Feed(ctx context.Context, opt *ListOptions, userId uint) ([]*Article, error)
	Favorite(ctx context.Context, slug string, userID uint) error
	UnFavorite(ctx context.Context, slug string, userID uint) error
	GetTags(ctx context.Context) ([]*Tag, error)
}
