package biz

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
)

type IpVote struct {
	IpVoteId, VoteId, PostId int
	Ip, Opts, TxtField       string
}

type IpVotes struct {
	Collection    []*IpVote
	NextPageToken string
}

type IpVoteRepo interface {
	ListIpVotes(ctx context.Context, parent string) (*IpVotes, error)
	GetIpVote(ctx context.Context, name string) (*IpVote, error)
	CreateIpVote(ctx context.Context, ipVote *IpVote) (*IpVote, error)
	UpdateIpVote(ctx context.Context, ipVote *IpVote) (*IpVote, error)
	DeleteIpVote(ctx context.Context, name string) (*emptypb.Empty, error)
}

type IpVoteUsecase struct {
	repo IpVoteRepo
}

func NewIpVoteUsecase(repo IpVoteRepo, logger log.Logger) *IpVoteUsecase {
	return &IpVoteUsecase{repo: repo}
}

func (ivu *IpVoteUsecase) ListIpVotes(ctx context.Context, parent string) (*IpVotes, error) {
	return ivu.repo.ListIpVotes(ctx, parent)
}

func (ivu *IpVoteUsecase) GetIpVote(ctx context.Context, name string) (*IpVote, error) {
	return ivu.repo.GetIpVote(ctx, name)
}

func (ivu *IpVoteUsecase) CreateIpVote(ctx context.Context, ipVote *IpVote) (*IpVote, error) {
	return ivu.repo.CreateIpVote(ctx, ipVote)
}

func (ivu *IpVoteUsecase) UpdateIpVote(ctx context.Context, ipVote *IpVote) (*IpVote, error) {
	return ivu.repo.UpdateIpVote(ctx, ipVote)
}

func (ivu *IpVoteUsecase) DeleteIpVote(ctx context.Context, name string) (*emptypb.Empty, error) {
	return ivu.repo.DeleteIpVote(ctx, name)
}
