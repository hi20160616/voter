package service

import (
	"context"
	"fmt"
	"log"

	pb "github.com/hi20160616/voter/api/voter/v1"
	"github.com/hi20160616/voter/internal/biz"
	"github.com/hi20160616/voter/internal/data"
	"github.com/hi20160616/voter/internal/data/db/mysql"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PostVoteService struct {
	pb.UnimplementedPostVotesAPIServer
	pc *biz.PostVoteUsecase
}

func NewPostVoteService() (*PostVoteService, error) {
	dbc, err := mysql.NewClient()
	if err != nil {
		return nil, err
	}

	db := &data.Data{DBClient: dbc}
	repo := data.NewPostVoteRepo(db, log.Default())
	postVoteUsecase := biz.NewPostVoteUsecase(repo, *log.Default())
	return &PostVoteService{pc: postVoteUsecase}, nil
}

func (pvs *PostVoteService) ListPostVotes(ctx context.Context, in *pb.ListPostVotesRequest) (*pb.ListPostVotesResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in ListPostVotes: \n%v\n", r)
		}
	}()
	bizpvs, err := pvs.pc.ListPostVotes(ctx, in.Parent)
	if err != nil {
		return nil, err
	}
	resp := []*pb.PostVote{}
	for _, p := range bizpvs.Collection {
		resp = append(resp, &pb.PostVote{
			PostVoteId: int32(p.PostVoteId),
			PostId:     int32(p.PostId),
			VoteId:     int32(p.VoteId),
		})
	}
	return &pb.ListPostVotesResponse{PostVotes: resp}, nil
}

func (pvs *PostVoteService) ListVidsByPid(ctx context.Context, in *pb.ListVidsByPidRequest) (*pb.ListVidsByPidResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in ListPostVotes: \n%v\n", r)
		}
	}()
	bizpvs, err := pvs.pc.ListVidsByPid(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	resp := &pb.ListVidsByPidResponse{Pid: int32(bizpvs.Pid)}
	for _, e := range bizpvs.Vids {
		resp.Vids = append(resp.Vids, int32(e))
	}
	return resp, nil
}

func (pvs *PostVoteService) GetByPidVid(ctx context.Context, in *pb.GetByPidVidRequest) (*pb.PostVote, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in GetPostVote: %s\n%v\n", in.Name, r)
		}
	}()
	bizpv, err := pvs.pc.GetByPidVid(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	return &pb.PostVote{
		PostVoteId: int32(bizpv.PostVoteId),
		PostId:     int32(bizpv.PostId),
		VoteId:     int32(bizpv.VoteId),
	}, nil
}

func (pvs *PostVoteService) GetPostVote(ctx context.Context, in *pb.GetPostVoteRequest) (*pb.PostVote, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in GetPostVote: %s\n%v\n", in.Name, r)
		}
	}()
	bizpv, err := pvs.pc.GetPostVote(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	return &pb.PostVote{
		PostVoteId: int32(bizpv.PostVoteId),
		PostId:     int32(bizpv.PostId),
		VoteId:     int32(bizpv.VoteId),
	}, nil
}

func (pvs *PostVoteService) CreatePostVote(ctx context.Context, in *pb.CreatePostVoteRequest) (*pb.PostVote, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdatePostVotes: \n%v\n", r)
		}
	}()
	_, err := pvs.pc.CreatePostVote(ctx, &biz.PostVote{
		PostId: int(in.PostVote.PostId),
		VoteId: int(in.PostVote.VoteId),
	})
	if err != nil {
		return nil, err
	}
	return &pb.PostVote{
		PostVoteId: in.PostVoteId,
		PostId:     in.PostVote.PostId,
		VoteId:     in.PostVote.VoteId,
	}, nil
}

func (pvs *PostVoteService) UpdatePostVote(ctx context.Context, in *pb.UpdatePostVoteRequest) (*pb.PostVote, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdatePostVote: \n%v\n", r)
		}
	}()
	p, err := pvs.pc.UpdatePostVote(ctx, &biz.PostVote{
		PostVoteId: int(in.PostVote.PostVoteId),
		PostId:     int(in.PostVote.PostId),
		VoteId:     int(in.PostVote.VoteId),
	})
	if err != nil {
		return nil, err
	}
	return &pb.PostVote{
		PostVoteId: int32(p.PostVoteId),
		PostId:     int32(p.PostId),
		VoteId:     int32(p.VoteId),
	}, nil
}

func (pvs *PostVoteService) DeletePostVote(ctx context.Context, in *pb.DeletePostVoteRequest) (*emptypb.Empty, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdatePostVotes: \n%v\n", r)
		}
	}()
	return pvs.pc.DeletePostVote(ctx, in.Name)
}
