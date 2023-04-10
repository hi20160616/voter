package biz

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
)

type IpPost struct {
	IpPostId, PostId int
	Ip               string
}

type IpPosts struct {
	Collection    []*IpPost
	NextPageToken string
}

type IpPostRepo interface {
	ListIpPosts(ctx context.Context, parent string) (*IpPosts, error)
	GetIpPost(ctx context.Context, name string) (*IpPost, error)
	CreateIpPost(ctx context.Context, ipPost *IpPost) (*IpPost, error)
	UpdateIpPost(ctx context.Context, ipPost *IpPost) (*IpPost, error)
	DeleteIpPost(ctx context.Context, name string) (*emptypb.Empty, error)
}

type IpPostUsecase struct {
	repo IpPostRepo
}

func NewIpPostUsecase(repo IpPostRepo, logger log.Logger) *IpPostUsecase {
	return &IpPostUsecase{repo: repo}
}

func (ivu *IpPostUsecase) ListIpPosts(ctx context.Context, parent string) (*IpPosts, error) {
	return ivu.repo.ListIpPosts(ctx, parent)
}

func (ivu *IpPostUsecase) GetIpPost(ctx context.Context, name string) (*IpPost, error) {
	return ivu.repo.GetIpPost(ctx, name)
}

func (ivu *IpPostUsecase) CreateIpPost(ctx context.Context, ipPost *IpPost) (*IpPost, error) {
	return ivu.repo.CreateIpPost(ctx, ipPost)
}

func (ivu *IpPostUsecase) UpdateIpPost(ctx context.Context, ipPost *IpPost) (*IpPost, error) {
	return ivu.repo.UpdateIpPost(ctx, ipPost)
}

func (ivu *IpPostUsecase) DeleteIpPost(ctx context.Context, name string) (*emptypb.Empty, error) {
	return ivu.repo.DeleteIpPost(ctx, name)
}
