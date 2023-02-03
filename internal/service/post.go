package service

import (
	"context"
	"fmt"

	"github.com/golang/glog"
	pb "github.com/hi20160616/voter/api/voter/v1"
	"github.com/hi20160616/voter/configs"
	"github.com/hi20160616/voter/internal/biz"
	"github.com/hi20160616/voter/internal/data/db/mysql"
	"google.golang.org/grpc"
)

type PostService struct {
	pb.UnimplementedPostsAPIServer
	pc *biz.PostUsecase
}

func NewPostService() (*PostService, error) {
	dbClient, err := mysql.NewClient()
	if err != nil {
		return nil, err
	}

	// db := &data.Data
}

func funHead(funName string, cfg *configs.Config) error {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorf("Recoved from %s: \n%v\n", funName, err)
		}
	}()
	if cfg == nil {
		return fmt.Errorf("%s: cfg is nil.", funName)
	}
	return nil
}

func ListPosts(ctx context.Context, in *pb.ListPostsRequest, cfg *configs.Config) (*pb.ListPostsResponse, error) {
	if err := funHead("ListPosts", cfg); err != nil {
		return nil, err
	}
	conn, err := grpc.Dial(cfg.API.GRPC.Addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewPostsAPIClient(conn)
	return c.ListPosts(ctx, in)
}

func GetPost(ctx context.Context, in *pb.GetPostRequest, cfg *configs.Config) (*pb.Post, error) {
	if err := funHead("GetPost", cfg); err != nil {
		return nil, err
	}
	conn, err := grpc.Dial(cfg.API.GRPC.Addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewPostsAPIClient(conn)
	return c.GetPost(ctx, in)
}

func SearchPosts(ctx context.Context, in *pb.SearchPostsRequest, cfg *configs.Config) (*pb.SearchPostsResponse, error) {
	if err := funHead("SearchPosts", cfg); err != nil {
		return nil, err
	}
	conn, err := grpc.Dial(cfg.API.GRPC.Addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewPostsAPIClient(conn)
	return c.SearchPosts(ctx, in)
}
