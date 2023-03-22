package biz

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
)

type PostVote struct {
	PostVoteId, PostId, VoteId int
}

type PostVotes struct {
	Collection    []*PostVote
	NextPageToken string
}

type PidVids struct {
	Pid  int
	Vids []int
}

type PostVoteRepo interface {
	ListPostVotes(ctx context.Context, parent string) (*PostVotes, error)
	GetPostVote(ctx context.Context, name string) (*PostVote, error)
	GetByPidVid(ctx context.Context, name string) (*PostVote, error)
	ListVidsByPid(ctx context.Context, name string) (*PidVids, error)
	CreatePostVote(ctx context.Context, postVote *PostVote) (*PostVote, error)
	UpdatePostVote(ctx context.Context, postVote *PostVote) (*PostVote, error)
	DeletePostVote(ctx context.Context, name string) (*emptypb.Empty, error)
}

type PostVoteUsecase struct {
	repo PostVoteRepo
}

func NewPostVoteUsecase(repo PostVoteRepo, logger log.Logger) *PostVoteUsecase {
	return &PostVoteUsecase{repo: repo}
}

func (pvu *PostVoteUsecase) ListPostVotes(ctx context.Context, parent string) (*PostVotes, error) {
	return pvu.repo.ListPostVotes(ctx, parent)
}

func (pvu *PostVoteUsecase) GetPostVote(ctx context.Context, name string) (*PostVote, error) {
	return pvu.repo.GetPostVote(ctx, name)
}

func (pvu *PostVoteUsecase) GetByPidVid(ctx context.Context, name string) (*PostVote, error) {
	return pvu.repo.GetByPidVid(ctx, name)
}

func (pvu *PostVoteUsecase) ListVidsByPid(ctx context.Context, name string) (*PidVids, error) {
	return pvu.repo.ListVidsByPid(ctx, name)
}

func (pvu *PostVoteUsecase) CreatePostVote(ctx context.Context, postVote *PostVote) (*PostVote, error) {
	return pvu.repo.CreatePostVote(ctx, postVote)
}

func (pvu *PostVoteUsecase) UpdatePostVote(ctx context.Context, postVote *PostVote) (*PostVote, error) {
	return pvu.repo.UpdatePostVote(ctx, postVote)
}

func (pvu *PostVoteUsecase) DeletePostVote(ctx context.Context, name string) (*emptypb.Empty, error) {
	return pvu.repo.DeletePostVote(ctx, name)
}
