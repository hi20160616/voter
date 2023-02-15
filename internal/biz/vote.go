package biz

import (
	"context"
	"log"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Vote struct {
	VoteId, IsRadio        int
	Title, Detail          string
	CreateTime, UpdateTime time.Time
}

type Votes struct {
	Collection    []*Vote
	NextPageToken string
}

type VoteRepo interface {
	ListVotes(ctx context.Context, parent string) (*Votes, error)
	GetVote(ctx context.Context, name string) (*Vote, error)
	SearchVotes(ctx context.Context, name string) (*Votes, error)
	CreateVote(ctx context.Context, vote *Vote) (*Vote, error)
	UpdateVote(ctx context.Context, vote *Vote) (*Vote, error)
	DeleteVote(ctx context.Context, name string) (*emptypb.Empty, error)
}

type VoteUsecase struct {
	repo VoteRepo
}

func NewVoteUsecase(repo VoteRepo, logger log.Logger) *VoteUsecase {
	return &VoteUsecase{repo: repo}
}

func (pu *VoteUsecase) ListVotes(ctx context.Context, parent string) (*Votes, error) {
	return pu.repo.ListVotes(ctx, parent)
}

func (pu *VoteUsecase) GetVote(ctx context.Context, name string) (*Vote, error) {
	return pu.repo.GetVote(ctx, name)
}

func (pu *VoteUsecase) SearchVotes(ctx context.Context, name string) (*Votes, error) {
	return pu.repo.SearchVotes(ctx, name)
}

func (pu *VoteUsecase) CreateVote(ctx context.Context, vote *Vote) (*Vote, error) {
	return pu.repo.CreateVote(ctx, vote)
}

func (pu *VoteUsecase) UpdateVote(ctx context.Context, vote *Vote) (*Vote, error) {
	return pu.repo.UpdateVote(ctx, vote)
}

func (pu *VoteUsecase) DeleteVote(ctx context.Context, name string) (*emptypb.Empty, error) {
	return pu.repo.DeleteVote(ctx, name)
}
