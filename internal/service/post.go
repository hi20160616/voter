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

type PostService struct {
	pb.UnimplementedPostsAPIServer
	pc *biz.PostUsecase
}

func NewPostService() (*PostService, error) {
	dbc, err := mysql.NewClient()
	if err != nil {
		return nil, err
	}

	db := &data.Data{DBClient: dbc}
	repo := data.NewPostRepo(db, log.Default())
	postUsecase := biz.NewPostUsecase(repo, *log.Default())
	return &PostService{pc: postUsecase}, nil
}

func (ps *PostService) ListPosts(ctx context.Context, in *pb.ListPostsRequest) (*pb.ListPostsResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in ListPosts: \n%v\n", r)
		}
	}()
	bizps, err := ps.pc.ListPosts(ctx, in.Parent)
	if err != nil {
		return nil, err
	}
	resp := []*pb.Post{}
	for _, p := range bizps.Collection {
		resp = append(resp, &pb.Post{
			PostId:     int32(p.PostId),
			Title:      p.Title,
			IsClosed:   int32(p.IsClosed),
			Detail:     p.Detail,
			CreateTime: timestamppb.New(p.CreateTime),
			UpdateTime: timestamppb.New(p.UpdateTime),
		})
	}
	return &pb.ListPostsResponse{Posts: resp}, nil
}

func (ps *PostService) GetPost(ctx context.Context, in *pb.GetPostRequest) (*pb.Post, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in GetPost: %s\n%v\n", in.Name, r)
		}
	}()
	bizp, err := ps.pc.GetPost(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	return &pb.Post{
		PostId:     int32(bizp.PostId),
		Title:      bizp.Title,
		IsClosed:   int32(bizp.IsClosed),
		Detail:     bizp.Detail,
		CreateTime: timestamppb.New(bizp.CreateTime),
		UpdateTime: timestamppb.New(bizp.UpdateTime),
	}, nil
}

func (ps *PostService) SearchPosts(ctx context.Context, in *pb.SearchPostsRequest) (*pb.SearchPostsResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in SearchPosts: \n%v\n", r)
		}
	}()
	bizps, err := ps.pc.SearchPosts(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	resp := &pb.SearchPostsResponse{}
	for _, e := range bizps.Collection {
		resp.Posts = append(resp.Posts, &pb.Post{
			PostId:     int32(e.PostId),
			Title:      e.Title,
			IsClosed:   int32(e.IsClosed),
			Detail:     e.Detail,
			CreateTime: timestamppb.New(e.CreateTime),
			UpdateTime: timestamppb.New(e.UpdateTime),
		})
	}
	return resp, nil
}

func (ps *PostService) CreatePost(ctx context.Context, in *pb.CreatePostRequest) (*pb.Post, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdatePosts: \n%v\n", r)
		}
	}()
	p, err := ps.pc.CreatePost(ctx, &biz.Post{
		Title:    in.Post.Title,
		IsClosed: int(in.Post.IsClosed),
		Detail:   in.Post.Detail,
	})
	if err != nil {
		return nil, err
	}
	return &pb.Post{
		Title:      p.Title,
		IsClosed:   int32(p.IsClosed),
		Detail:     p.Detail,
		UpdateTime: timestamppb.New(p.UpdateTime),
		CreateTime: timestamppb.New(p.CreateTime),
	}, nil
}

func (ps *PostService) UpdatePost(ctx context.Context, in *pb.UpdatePostRequest) (*pb.Post, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdatePost: \n%v\n", r)
		}
	}()
	p, err := ps.pc.UpdatePost(ctx, &biz.Post{
		PostId:   int(in.Post.PostId),
		Title:    in.Post.Title,
		IsClosed: int(in.Post.IsClosed),
		Detail:   in.Post.Detail,
	})
	if err != nil {
		return nil, err
	}
	return &pb.Post{
		PostId:     int32(p.PostId),
		Title:      p.Title,
		IsClosed:   int32(p.IsClosed),
		Detail:     p.Detail,
		UpdateTime: timestamppb.New(p.UpdateTime),
		CreateTime: timestamppb.New(p.CreateTime),
	}, nil
}

func (ps *PostService) DeletePost(ctx context.Context, in *pb.DeletePostRequest) (*emptypb.Empty, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdatePosts: \n%v\n", r)
		}
	}()
	return ps.pc.DeletePost(ctx, in.Name)
}
