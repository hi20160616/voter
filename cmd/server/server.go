package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"

	"github.com/golang/glog"
	"github.com/hi20160616/voter/configs"
	"github.com/hi20160616/voter/internal/server/web"

	theGateway "github.com/hi20160616/voter/internal/server/gateway"
	theGRPC "github.com/hi20160616/voter/internal/server/grpc"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	g, ctx := errgroup.WithContext(ctx)
	cfg := configs.NewConfig("voter")

	// Http server
	g.Go(func() error {
		return web.Run(ctx, cfg)
	})

	// gRPC
	g.Go(func() error { return theGRPC.Run(ctx, "tcp", cfg.API.GRPC.Addr) })

	// gRPC-gateway
	opts := theGateway.Options{
		Addr: cfg.API.HTTP.Addr,
		GRPCServer: theGateway.Endpoint{
			Network: cfg.API.GRPC.Network,
			Addr:    cfg.API.GRPC.Addr,
		},
	}
	g.Go(func() error { return theGateway.Run(ctx, opts) })

	// Graceful stop
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	g.Go(func() error {
		select {
		case sig := <-sigs:
			fmt.Println()
			glog.Infof("signal caught: %s, ready to quit...", sig.String())
			cancel()
		case <-ctx.Done():
			return ctx.Err()
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		if !errors.Is(err, context.Canceled) {
			glog.Errorf("not canceled by context: %s", err)
		} else {
			glog.Info(err)
		}
	}
}
