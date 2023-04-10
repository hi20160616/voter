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

type IpPostService struct {
	pb.UnimplementedIpPostsAPIServer
	pc *biz.IpPostUsecase
}

func NewIpPostService() (*IpPostService, error) {
	dbc, err := mysql.NewClient()
	if err != nil {
		return nil, err
	}

	db := &data.Data{DBClient: dbc}
	repo := data.NewIpPostRepo(db, log.Default())
	ipPostUsecase := biz.NewIpPostUsecase(repo, *log.Default())
	return &IpPostService{pc: ipPostUsecase}, nil
}

func (ips *IpPostService) ListIpPosts(ctx context.Context, in *pb.ListIpPostsRequest) (*pb.ListIpPostsResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in ListIpPosts: \n%v\n", r)
		}
	}()
	bizips, err := ips.pc.ListIpPosts(ctx, in.Parent)
	if err != nil {
		return nil, err
	}
	resp := []*pb.IpPost{}
	for _, p := range bizips.Collection {
		resp = append(resp, &pb.IpPost{
			IpPostId: int32(p.IpPostId),
			Ip:       p.Ip,
			PostId:   int32(p.PostId),
		})
	}
	return &pb.ListIpPostsResponse{IpPosts: resp}, nil
}

func (ips *IpPostService) GetIpPost(ctx context.Context, in *pb.GetIpPostRequest) (*pb.IpPost, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in GetIpPost: %s\n%v\n", in.Name, r)
		}
	}()
	bizip, err := ips.pc.GetIpPost(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	return &pb.IpPost{
		IpPostId: int32(bizip.IpPostId),
		Ip:       bizip.Ip,
		PostId:   int32(bizip.PostId),
	}, nil
}

func (ips *IpPostService) CreateIpPost(ctx context.Context, in *pb.CreateIpPostRequest) (*pb.IpPost, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdateIpPosts: \n%v\n", r)
		}
	}()
	iv, err := ips.pc.CreateIpPost(ctx, &biz.IpPost{
		Ip:     in.IpPost.Ip,
		PostId: int(in.IpPost.PostId),
	})
	if err != nil {
		return nil, err
	}
	return &pb.IpPost{
		IpPostId: int32(iv.IpPostId),
		Ip:       iv.Ip,
		PostId:   int32(iv.PostId),
	}, nil
}

func (ips *IpPostService) UpdateIpPost(ctx context.Context, in *pb.UpdateIpPostRequest) (*pb.IpPost, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdateIpPost: \n%v\n", r)
		}
	}()
	p, err := ips.pc.UpdateIpPost(ctx, &biz.IpPost{
		IpPostId: int(in.IpPost.IpPostId),
		Ip:       in.IpPost.Ip,
		PostId:   int(in.IpPost.PostId),
	})
	if err != nil {
		return nil, err
	}
	return &pb.IpPost{
		IpPostId: int32(p.IpPostId),
		Ip:       p.Ip,
		PostId:   int32(p.PostId),
	}, nil
}

func (ips *IpPostService) DeleteIpPost(ctx context.Context, in *pb.DeleteIpPostRequest) (*emptypb.Empty, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdateIpPosts: \n%v\n", r)
		}
	}()
	return ips.pc.DeleteIpPost(ctx, in.Name)
}
