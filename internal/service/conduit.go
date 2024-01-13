package service

import (
	"context"

	pb "kratos-realworld/api/conduit/v1"
	"kratos-realworld/internal/biz"

	"github.com/golang/protobuf/ptypes/empty"
)

type ConduitService struct {
	pb.UnimplementedConduitServer

	uc *biz.ConduitUseCase
}

func NewConduitService(uc *biz.ConduitUseCase) *ConduitService {
	return &ConduitService{uc: uc}
}

func (s *ConduitService) Authenticate(ctx context.Context, req *pb.AuthenticateRequest) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}

func (s *ConduitService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}

func (s *ConduitService) GetCurrentUser(ctx context.Context, req *empty.Empty) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}

func (s *ConduitService) UpdateUser(ctx context.Context, req *pb.UserReply_User) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}

func (s *ConduitService) GetProfile(ctx context.Context, req *pb.ProfileRequest) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}

func (s *ConduitService) FollowUser(ctx context.Context, req *pb.ProfileRequest) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}

func (s *ConduitService) UnfollowUser(ctx context.Context, req *pb.ProfileRequest) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}

func (s *ConduitService) ListArticles(ctx context.Context, req *pb.ListArticlesRequest) (*pb.ArticlesReply, error) {
	return &pb.ArticlesReply{}, nil
}

func (s *ConduitService) FeedArticles(ctx context.Context, req *pb.ListArticlesRequest) (*pb.ArticlesReply, error) {
	return &pb.ArticlesReply{}, nil
}

func (s *ConduitService) GetArticle(ctx context.Context, req *pb.ArticleRequest) (*pb.ArticleReply, error) {
	return &pb.ArticleReply{}, nil
}

func (s *ConduitService) CreateArticle(ctx context.Context, req *pb.ArticleReply_Article) (*pb.ArticleReply, error) {
	return &pb.ArticleReply{}, nil
}

func (s *ConduitService) UpdateArticle(ctx context.Context, req *pb.ArticleReply_Article) (*pb.ArticleReply, error) {
	return &pb.ArticleReply{}, nil
}

func (s *ConduitService) DeleteArticle(ctx context.Context, req *pb.ArticleRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *ConduitService) AddComment(ctx context.Context, req *pb.AddCommentRequest) (*pb.CommentReply, error) {
	return &pb.CommentReply{}, nil
}

func (s *ConduitService) GetComments(ctx context.Context, req *pb.GetCommentsRequest) (*pb.CommentsReply, error) {
	return &pb.CommentsReply{}, nil
}

func (s *ConduitService) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *ConduitService) FavoriteArticle(ctx context.Context, req *pb.ArticleRequest) (*pb.ArticleReply, error) {
	return &pb.ArticleReply{}, nil
}

func (s *ConduitService) UnfavoriteArticle(ctx context.Context, req *pb.ArticleRequest) (*pb.ArticleReply, error) {
	return &pb.ArticleReply{}, nil
}

func (s *ConduitService) ListTags(ctx context.Context, req *empty.Empty) (*pb.TagsReply, error) {
	return &pb.TagsReply{}, nil
}
