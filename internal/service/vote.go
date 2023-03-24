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
	vc *biz.VoteUsecase
}

func NewVoteService() (*VoteService, error) {
	dbc, err := mysql.NewClient()
	if err != nil {
		return nil, err
	}

	db := &data.Data{DBClient: dbc}
	repo := data.NewVoteRepo(db, log.Default())
	voteUsecase := biz.NewVoteUsecase(repo, *log.Default())
	return &VoteService{vc: voteUsecase}, nil
}

func (vs *VoteService) ListVotes(ctx context.Context, in *pb.ListVotesRequest) (*pb.ListVotesResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in ListVotes: \n%v\n", r)
		}
	}()
	bizvs, err := vs.vc.ListVotes(ctx, in.Parent)
	if err != nil {
		return nil, err
	}
	resp := []*pb.Vote{}
	for _, u := range bizvs.Collection {
		resp = append(resp, &pb.Vote{
			VoteId:     int32(u.VoteId),
			Title:      u.Title,
			A:          u.A,
			B:          u.B,
			C:          u.C,
			D:          u.D,
			E:          u.E,
			F:          u.F,
			G:          u.G,
			H:          u.H,
			IsRadio:    int32(u.IsRadio),
			CreateTime: timestamppb.New(u.CreateTime),
			UpdateTime: timestamppb.New(u.UpdateTime),
		})
	}
	return &pb.ListVotesResponse{Votes: resp}, nil
}

func (vs *VoteService) GetVote(ctx context.Context, in *pb.GetVoteRequest) (*pb.Vote, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in GetVote: %s\n%v\n", in.Name, r)
		}
	}()
	bizv, err := vs.vc.GetVote(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	return &pb.Vote{
		VoteId:      int32(bizv.VoteId),
		Title:       bizv.Title,
		IsRadio:     int32(bizv.IsRadio),
		A:           bizv.A,
		B:           bizv.B,
		C:           bizv.C,
		D:           bizv.D,
		E:           bizv.E,
		F:           bizv.F,
		G:           bizv.G,
		H:           bizv.H,
		HasTxtField: int32(bizv.HasTxtField),
		CreateTime:  timestamppb.New(bizv.CreateTime),
		UpdateTime:  timestamppb.New(bizv.UpdateTime),
	}, nil
}

func (us *VoteService) SearchVotes(ctx context.Context, in *pb.SearchVotesRequest) (*pb.SearchVotesResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in SearchVotes: \n%v\n", r)
		}
	}()
	bizvs, err := us.vc.SearchVotes(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	resp := &pb.SearchVotesResponse{}
	for _, e := range bizvs.Collection {
		resp.Votes = append(resp.Votes, &pb.Vote{
			VoteId:      int32(e.VoteId),
			Title:       e.Title,
			IsRadio:     int32(e.IsRadio),
			A:           e.A,
			B:           e.B,
			C:           e.C,
			D:           e.D,
			E:           e.E,
			F:           e.F,
			G:           e.G,
			H:           e.H,
			HasTxtField: int32(e.HasTxtField),
			CreateTime:  timestamppb.New(e.CreateTime),
			UpdateTime:  timestamppb.New(e.UpdateTime),
		})
	}
	return resp, nil
}

func (vs *VoteService) CreateVote(ctx context.Context, in *pb.CreateVoteRequest) (*pb.Vote, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdateVotes: \n%v\n", r)
		}
	}()
	v, err := vs.vc.CreateVote(ctx, &biz.Vote{
		VoteId:      int(in.Vote.VoteId),
		Title:       in.Vote.Title,
		IsRadio:     int(in.Vote.IsRadio),
		A:           in.Vote.A,
		B:           in.Vote.B,
		C:           in.Vote.C,
		D:           in.Vote.D,
		E:           in.Vote.E,
		F:           in.Vote.F,
		G:           in.Vote.G,
		H:           in.Vote.H,
		HasTxtField: int(in.Vote.HasTxtField),
	})
	if err != nil {
		return nil, err
	}
	return &pb.Vote{
		VoteId:      int32(v.VoteId),
		Title:       v.Title,
		IsRadio:     int32(v.IsRadio),
		A:           v.A,
		B:           v.B,
		C:           v.C,
		D:           v.D,
		E:           v.E,
		F:           v.F,
		G:           v.G,
		H:           v.H,
		HasTxtField: int32(v.HasTxtField),
		CreateTime:  timestamppb.New(v.CreateTime),
		UpdateTime:  timestamppb.New(v.UpdateTime),
	}, nil
}

func (vs *VoteService) UpdateVote(ctx context.Context, in *pb.UpdateVoteRequest) (*pb.Vote, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdateVotes: \n%v\n", r)
		}
	}()
	v, err := vs.vc.UpdateVote(ctx, &biz.Vote{
		VoteId:      int(in.Vote.VoteId),
		Title:       in.Vote.Title,
		IsRadio:     int(in.Vote.IsRadio),
		A:           in.Vote.A,
		B:           in.Vote.B,
		C:           in.Vote.C,
		D:           in.Vote.D,
		E:           in.Vote.E,
		F:           in.Vote.F,
		G:           in.Vote.G,
		H:           in.Vote.H,
		HasTxtField: int(in.Vote.HasTxtField),
	})
	if err != nil {
		return nil, err
	}
	return &pb.Vote{
		VoteId:      int32(v.VoteId),
		Title:       v.Title,
		IsRadio:     int32(v.IsRadio),
		A:           v.A,
		B:           v.B,
		C:           v.C,
		D:           v.D,
		E:           v.E,
		F:           v.F,
		G:           v.G,
		H:           v.H,
		HasTxtField: int32(v.HasTxtField),
		CreateTime:  timestamppb.New(v.CreateTime),
		UpdateTime:  timestamppb.New(v.UpdateTime),
	}, nil
}

func (vs *VoteService) DeleteVote(ctx context.Context, in *pb.DeleteVoteRequest) (*emptypb.Empty, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in DeleteVotes: \n%v\n", r)
		}
	}()
	return vs.vc.DeleteVote(ctx, in.Name)
}
