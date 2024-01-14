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

// Create Article
//
//	(1) Authentication required, returns the created Article
//	(2) Required fields: title, description, body
//	(3) Optional fields: tagList as []string
func (uc *ConduitUseCase) CreateArticle(ctx context.Context, article *Article) (*Article, error) {
	return uc.ar.Create(ctx, article)
}

// Get Article
//
//	(1) No authentication required, will return single article
func (uc *ConduitUseCase) GetArticle(ctx context.Context, slug string) (*Article, error) {
	return uc.ar.Get(ctx, slug)
}

// Update Article
//
//	(1) Authentication required, returns the updated Article
//	(2) Optional fields: title, description, body	
//	(3) Slug also gets updated when the title is changed
func (uc *ConduitUseCase) UpdateArticle(ctx context.Context, slug string, article *Article) (*Article, error) {
	return uc.ar.Update(ctx, slug, article)
}

// Delete Article
//
//	(1) Authentication required
func (uc *ConduitUseCase) DeleteArticle(ctx context.Context, slug string) error {
	return uc.ar.Delete(ctx, slug)
}

// List articles
//
//	(1) Returns most recent articles globally by default, provide tag, author or favorited query parameter to filter results
//	(2) Limit number of articles (default is 20):
//	(3) Offset/skip number of articles (default is 0):
//	(4) Authentication optional, will return multiple articles, ordered by most recent first
func (uc *ConduitUseCase) ListArticle(ctx context.Context, opt *ListOptions) ([]*Article, error) {
	return uc.ar.List(ctx, opt)
}

// Feed articles
//	
//	(1) Can also take limit and offset query parameters like List Articles
//	(2) Authentication required, will return multiple articles created by followed users, ordered by most recent first	
func (uc *ConduitUseCase) FeedArticle(ctx context.Context, opt *ListOptions, userId uint) ([]*Article, error) {
	return uc.ar.Feed(ctx, opt, userId)
}

// Favorite Article
//
//	(1) Authentication required, returns the Article
//	(2) No additional parameters required
func (uc *ConduitUseCase) FavoriteArticle(ctx context.Context, slug string, userID uint) error {
	return uc.ar.Favorite(ctx, slug, userID)
}

// Unfavorite Article
//
//	(1) Authentication required, returns the Article
//	(2) No additional parameters required
func (uc *ConduitUseCase) UnFavoriteArticle(ctx context.Context, slug string, userID uint) error {
	return uc.ar.UnFavorite(ctx, slug, userID)
}

// Get Tags
//	
//	(1) No authentication required, returns a List of Tags
func (uc *ConduitUseCase) GetArticleTags(ctx context.Context) ([]*Tag, error) {
	return uc.ar.GetTags(ctx)
}
