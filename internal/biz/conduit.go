package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

// ConduitUseCase is a Conduit UseCase.
type ConduitUseCase struct {
	ur  UserRepo
	ar  ArticleRepo
	cr  CommentRepo
	pr  ProfileRepo
	log *log.Helper
}

// NewConduitUseCase new a Conduit UseCase.
func NewConduitUseCase(
	ur UserRepo,
	ar ArticleRepo,
	cr CommentRepo,
	pr ProfileRepo,
	logger log.Logger) *ConduitUseCase {
	return &ConduitUseCase{
		ur:  ur,
		ar:  ar,
		cr:  cr,
		pr:  pr,
		log: log.NewHelper(logger)}
}
