package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

// ConduitUseCase is a Conduit UseCase.
type ConduitUseCase struct {
	ar  ArticleRepo
	cr  CommentRepo
	ur  UserRepo
	pr  ProfileRepo
	log *log.Helper
}

// NewConduitUseCase new a Conduit UseCase.
func NewConduitUseCase(
	ar ArticleRepo,
	cr CommentRepo,
	ur UserRepo,
	pr ProfileRepo,
	logger log.Logger) *ConduitUseCase {
	return &ConduitUseCase{
		ar:  ar,
		cr:  cr,
		ur:  ur,
		pr:  pr,
		log: log.NewHelper(logger)}
}
