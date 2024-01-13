package data

import (
	v1 "kratos-realworld/api/conduit/v1"
	"kratos-realworld/internal/conf"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var (
	// ErrNotImplemented is comment not found.
	ErrNotImplemented = errors.NotFound(v1.ErrorReason_NOT_IMPLEMENTED.String(), "not implemented")
	// ProviderSet is data providers.
	ProviderSet = wire.NewSet(NewData, NewArticleRepo, NewUserRepo, NewCommentRepo, NewProfileRepo)
)

// Data .
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}
