package grpc

import (
	"context"
	"net"

	"github.com/golang/glog"
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

	// ps, err := service.NewPost
	return nil
}
