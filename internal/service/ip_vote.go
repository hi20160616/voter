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

type IpVoteService struct {
	pb.UnimplementedIpVotesAPIServer
	pc *biz.IpVoteUsecase
}

func NewIpVoteService() (*IpVoteService, error) {
	dbc, err := mysql.NewClient()
	if err != nil {
		return nil, err
	}

	db := &data.Data{DBClient: dbc}
	repo := data.NewIpVoteRepo(db, log.Default())
	ipVoteUsecase := biz.NewIpVoteUsecase(repo, *log.Default())
	return &IpVoteService{pc: ipVoteUsecase}, nil
}

func (ivs *IpVoteService) ListIpVotes(ctx context.Context, in *pb.ListIpVotesRequest) (*pb.ListIpVotesResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in ListIpVotes: \n%v\n", r)
		}
	}()
	bizivs, err := ivs.pc.ListIpVotes(ctx, in.Parent)
	if err != nil {
		return nil, err
	}
	resp := []*pb.IpVote{}
	for _, p := range bizivs.Collection {
		resp = append(resp, &pb.IpVote{
			IpVoteId: int32(p.IpVoteId),
			Ip:       p.Ip,
			VoteId:   int32(p.VoteId),
			Opts:     p.Opts,
			TxtField: p.TxtField,
		})
	}
	return &pb.ListIpVotesResponse{IpVotes: resp}, nil
}

func (ivs *IpVoteService) GetIpVote(ctx context.Context, in *pb.GetIpVoteRequest) (*pb.IpVote, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in GetIpVote: %s\n%v\n", in.Name, r)
		}
	}()
	bizpv, err := ivs.pc.GetIpVote(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	return &pb.IpVote{
		IpVoteId: int32(bizpv.IpVoteId),
		Ip:       bizpv.Ip,
		VoteId:   int32(bizpv.VoteId),
		Opts:     bizpv.Opts,
		TxtField: bizpv.TxtField,
	}, nil
}

func (ivs *IpVoteService) CreateIpVote(ctx context.Context, in *pb.CreateIpVoteRequest) (*pb.IpVote, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdateIpVotes: \n%v\n", r)
		}
	}()
	iv, err := ivs.pc.CreateIpVote(ctx, &biz.IpVote{
		Ip:       in.IpVote.Ip,
		VoteId:   int(in.IpVote.VoteId),
		Opts:     in.IpVote.Opts,
		TxtField: in.IpVote.TxtField,
	})
	if err != nil {
		return nil, err
	}
	return &pb.IpVote{
		IpVoteId: int32(iv.IpVoteId),
		Ip:       iv.Ip,
		VoteId:   int32(iv.VoteId),
		Opts:     iv.Opts,
		TxtField: iv.TxtField,
	}, nil
}

func (ivs *IpVoteService) UpdateIpVote(ctx context.Context, in *pb.UpdateIpVoteRequest) (*pb.IpVote, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdateIpVote: \n%v\n", r)
		}
	}()
	p, err := ivs.pc.UpdateIpVote(ctx, &biz.IpVote{
		IpVoteId: int(in.IpVote.IpVoteId),
		Ip:       in.IpVote.Ip,
		VoteId:   int(in.IpVote.VoteId),
		Opts:     in.IpVote.Opts,
		TxtField: in.IpVote.TxtField,
	})
	if err != nil {
		return nil, err
	}
	return &pb.IpVote{
		IpVoteId: int32(p.IpVoteId),
		Ip:       p.Ip,
		VoteId:   int32(p.VoteId),
		Opts:     p.Opts,
		TxtField: p.TxtField,
	}, nil
}

func (ivs *IpVoteService) DeleteIpVote(ctx context.Context, in *pb.DeleteIpVoteRequest) (*emptypb.Empty, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdateIpVotes: \n%v\n", r)
		}
	}()
	return ivs.pc.DeleteIpVote(ctx, in.Name)
}
