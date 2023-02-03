package biz

import (
	"context"
	"log"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Post struct {
	IsOpen                 int
	Title, Detail          string
	CreateTime, UpdateTime time.Time
}

type Posts struct {
	Collention    []*Post
	NextPageToken string
}

type PostRepo interface {
	ListPosts(ctx context.Context, parent string) (*Posts, error)
	GetPost(ctx context.Context, name string) (*Post, error)
	SearchPosts(ctx context.Context, name string) (*Posts, error)
	CreatePost(ctx context.Context, post *Post) (*Post, error)
	UpdatePost(ctx context.Context, post *Post) (*Post, error)
	DeletePost(ctx context.Context, name string) (*emptypb.Empty, error)
}

type PostUsecase struct {
	repo PostRepo
}

func NewPostUsecase(repo PostRepo, logger log.Logger) *PostUsecase {
	return &PostUsecase{repo: repo}
}

func (pu *PostUsecase) ListPosts(ctx context.Context, parent string) (*Posts, error) {
	return pu.repo.ListPosts(ctx, parent)
}

func (pu *PostUsecase) GetPost(ctx context.Context, name string) (*Post, error) {
	return pu.repo.GetPost(ctx, name)
}

func (pu *PostUsecase) SearchPosts(ctx context.Context, name string) (*Posts, error) {
	return pu.repo.SearchPosts(ctx, name)
}

func (pu *PostUsecase) CreatePost(ctx context.Context, post *Post) (*Post, error) {
	return pu.repo.CreatePost(ctx, post)
}

func (pu *PostUsecase) UpdatePost(ctx context.Context, post *Post) (*Post, error) {
	return pu.repo.UpdatePost(ctx, post)
}

func (pu *PostUsecase) DeletePost(ctx context.Context, name string) (*emptypb.Empty, error) {
	return pu.repo.DeletePost(ctx, name)
}
