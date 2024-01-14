package service

import (
	"context"

	pb "kratos-realworld/api/conduit/v1"
	"kratos-realworld/internal/biz"
)

type ConduitService struct {
	pb.UnimplementedConduitServer

	uc *biz.ConduitUseCase
}

func NewConduitService(uc *biz.ConduitUseCase) *ConduitService {
	return &ConduitService{uc: uc}
}

func (s *ConduitService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
func (s *ConduitService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
func (s *ConduitService) GetCurrentUser(ctx context.Context, req *pb.GetCurrentUserRequest) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
func (s *ConduitService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
func (s *ConduitService) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}
func (s *ConduitService) FollowUser(ctx context.Context, req *pb.FollowUserRequest) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}
func (s *ConduitService) UnfollowUser(ctx context.Context, req *pb.UnfollowUserRequest) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}
func (s *ConduitService) ListArticles(ctx context.Context, req *pb.ListArticlesRequest) (*pb.ArticlesReply, error) {
	return &pb.ArticlesReply{}, nil
}
func (s *ConduitService) FeedArticles(ctx context.Context, req *pb.FeedArticlesRequest) (*pb.ArticlesReply, error) {
	return &pb.ArticlesReply{}, nil
}
func (s *ConduitService) GetArticle(ctx context.Context, req *pb.GetArticleRequest) (*pb.ArticleReply, error) {
	return &pb.ArticleReply{}, nil
}
func (s *ConduitService) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.ArticleReply, error) {
	return &pb.ArticleReply{}, nil
}
func (s *ConduitService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.ArticleReply, error) {
	return &pb.ArticleReply{}, nil
}
func (s *ConduitService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.EmptyReply, error) {
	return &pb.EmptyReply{}, nil
}
func (s *ConduitService) AddComment(ctx context.Context, req *pb.AddCommentRequest) (*pb.CommentReply, error) {
	return &pb.CommentReply{}, nil
}
func (s *ConduitService) GetComments(ctx context.Context, req *pb.GetCommentsRequest) (*pb.CommentsReply, error) {
	return &pb.CommentsReply{}, nil
}
func (s *ConduitService) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.EmptyReply, error) {
	return &pb.EmptyReply{}, nil
}
func (s *ConduitService) FavoriteArticle(ctx context.Context, req *pb.FavoriteArticleRequest) (*pb.ArticleReply, error) {
	return &pb.ArticleReply{}, nil
}
func (s *ConduitService) UnfavoriteArticle(ctx context.Context, req *pb.UnfavoriteArticleRequest) (*pb.ArticleReply, error) {
	return &pb.ArticleReply{}, nil
}
func (s *ConduitService) ListTags(ctx context.Context, req *pb.ListTagsRequest) (*pb.TagsReply, error) {
	return &pb.TagsReply{}, nil
}
