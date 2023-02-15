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
	"google.golang.org/protobuf/types/known/timestamppb"
)

type VoteService struct {
	pb.UnimplementedVotesAPIServer
	uc *biz.VoteUsecase
}

func NewVoteService() (*VoteService, error) {
	dbc, err := mysql.NewClient()
	if err != nil {
		return nil, err
	}

	db := &data.Data{DBClient: dbc}
	repo := data.NewVoteRepo(db, log.Default())
	voteUsecase := biz.NewVoteUsecase(repo, *log.Default())
	return &VoteService{uc: voteUsecase}, nil
}

func (vs *VoteService) ListVotes(ctx context.Context, in *pb.ListVotesRequest) (*pb.ListVotesResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in ListVotes: \n%v\n", r)
		}
	}()
	bizvs, err := vs.uc.ListVotes(ctx, in.Parent)
	if err != nil {
		return nil, err
	}
	resp := []*pb.Vote{}
	for _, u := range bizvs.Collection {
		resp = append(resp, &pb.Vote{
			VoteId:     int32(u.VoteId),
			Title:      u.Title,
			IsRadio:    int32(u.IsRadio),
			Detail:     u.Detail,
			CreateTime: timestamppb.New(u.CreateTime),
			UpdateTime: timestamppb.New(u.UpdateTime),
		})
	}
	return &pb.ListVotesResponse{Votes: resp}, nil
}

func (us *VoteService) GetVote(ctx context.Context, in *pb.GetVoteRequest) (*pb.Vote, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in GetVote: %s\n%v\n", in.Name, r)
		}
	}()
	bizu, err := us.uc.GetVote(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	return &pb.Vote{
		VoteId:     int32(bizu.VoteId),
		Title:      bizu.Title,
		IsRadio:    int32(bizu.IsRadio),
		Detail:     bizu.Detail,
		CreateTime: timestamppb.New(bizu.CreateTime),
		UpdateTime: timestamppb.New(bizu.UpdateTime),
	}, nil
}

func (us *VoteService) SearchVotes(ctx context.Context, in *pb.SearchVotesRequest) (*pb.SearchVotesResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in SearchVotes: \n%v\n", r)
		}
	}()
	bizvs, err := us.uc.SearchVotes(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	resp := &pb.SearchVotesResponse{}
	for _, e := range bizvs.Collection {
		resp.Votes = append(resp.Votes, &pb.Vote{
			VoteId:     int32(e.VoteId),
			Title:      e.Title,
			IsRadio:    int32(e.IsRadio),
			Detail:     e.Detail,
			CreateTime: timestamppb.New(e.CreateTime),
			UpdateTime: timestamppb.New(e.UpdateTime),
		})
	}
	return resp, nil
}

func (us *VoteService) CreateVote(ctx context.Context, in *pb.CreateVoteRequest) (*pb.Vote, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdateVotes: \n%v\n", r)
		}
	}()
	u, err := us.uc.CreateVote(ctx, &biz.Vote{
		VoteId:  int(in.Vote.VoteId),
		Title:   in.Vote.Title,
		IsRadio: int(in.Vote.IsRadio),
		Detail:  in.Vote.Detail,
	})
	if err != nil {
		return nil, err
	}
	return &pb.Vote{
		VoteId:     int32(u.VoteId),
		Title:      u.Title,
		IsRadio:    int32(u.IsRadio),
		Detail:     u.Detail,
		CreateTime: timestamppb.New(u.CreateTime),
		UpdateTime: timestamppb.New(u.UpdateTime),
	}, nil
}

func (us *VoteService) UpdateVote(ctx context.Context, in *pb.UpdateVoteRequest) (*pb.Vote, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdateVotes: \n%v\n", r)
		}
	}()
	u, err := us.uc.UpdateVote(ctx, &biz.Vote{
		VoteId:  int(in.Vote.VoteId),
		Title:   in.Vote.Title,
		IsRadio: int(in.Vote.IsRadio),
		Detail:  in.Vote.Detail,
	})
	if err != nil {
		return nil, err
	}
	return &pb.Vote{
		VoteId:     int32(u.VoteId),
		Title:      u.Title,
		IsRadio:    int32(u.IsRadio),
		UpdateTime: timestamppb.New(u.UpdateTime),
		CreateTime: timestamppb.New(u.CreateTime),
	}, nil
}

func (us *VoteService) DeleteVote(ctx context.Context, in *pb.DeleteVoteRequest) (*emptypb.Empty, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in DeleteVotes: \n%v\n", r)
		}
	}()
	return us.uc.DeleteVote(ctx, in.Name)
}
