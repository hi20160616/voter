package grpc

import (
	"context"
	"net"

	"github.com/golang/glog"
	pb "github.com/hi20160616/voter/api/voter/v1"
	"github.com/hi20160616/voter/internal/service"
	"google.golang.org/grpc"
)

func Run(ctx context.Context, network, address string) error {
	l, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	defer func() {
		if err := l.Close(); err != nil {
			glog.Errorf("Failed to close %s %s: %v", network, address, err)
		}
	}()

	ps, err := service.NewPostService()
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterPostsAPIServer(s, ps)
	go func() {
		defer s.GracefulStop()
		<-ctx.Done()
	}()

	glog.Infof("gRPC starting listening at %s", address)
	return s.Serve(l)
}
