package gateway

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/golang/glog"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/hi20160616/voter/api/voter/v1"
	"google.golang.org/grpc"
)

type Options struct {
	Addr       string   // gateway address to listen
	GRPCServer Endpoint // gRPC service endpoint
	// Mux        []gwruntime.ServeMuxOption
}

type Endpoint struct {
	Network, Addr string
}

// Run starts a HTTP server and blocks while running if successful.
// The server will be shutdown when "ctx" is canceled.
func Run(ctx context.Context, opts Options) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Conn
	conn, err := dial(ctx, opts.GRPCServer.Network, opts.GRPCServer.Addr)
	if err != nil {
		return err
	}
	// Conn Close
	go func() {
		<-ctx.Done()
		if err := conn.Close(); err != nil {
			glog.Errorf("Failed to close a client connection to the gRPC server: %v", err)
		}
	}()
	// New Gateway mux
	gwServer, err := NewGatewayServer(ctx, conn, opts)
	if err != nil {
		return err
	}
	// Gateway Graceful Stop
	go func() {
		<-ctx.Done()
		glog.Infof("Shutting down the grpc-gateway http server")
		if err := gwServer.Shutdown(ctx); err != nil {
			glog.Errorf("Failed to shutdown grpc-gateway http server: %v", err)
		}
	}()
	// Gateway Start
	glog.Infof("gRPC gateway starting listening at %s", opts.Addr)
	if err := gwServer.ListenAndServe(); err != http.ErrServerClosed {
		glog.Errorf("Failed to listen and serve: %v", err)
		return err
	}
	return nil
}

func NewGatewayServer(ctx context.Context, conn *grpc.ClientConn, opts Options) (*http.Server, error) {
	mux := gwruntime.NewServeMux()
	for _, f := range []func(context.Context, *gwruntime.ServeMux, *grpc.ClientConn) error{
		pb.RegisterPostsAPIHandler,
		// pb.RegisterTagsAPIHandler,
	} {
		if err := f(ctx, mux, conn); err != nil {
			return nil, err
		}

	}
	return &http.Server{
		Addr:    opts.Addr,
		Handler: mux,
	}, nil
}

func dial(ctx context.Context, network, addr string) (*grpc.ClientConn, error) {
	switch strings.ToLower(network) {
	case "tcp":
		return dialTCP(ctx, addr)
	case "unix":
		return dialUnix(ctx, addr)
	default:
		return nil, fmt.Errorf("unsupported network type: %s", network)
	}
}

func dialTCP(ctx context.Context, addr string) (*grpc.ClientConn, error) {
	return grpc.DialContext(ctx, addr, grpc.WithInsecure())
}

func dialUnix(ctx context.Context, addr string) (*grpc.ClientConn, error) {
	d := func(ctx context.Context, addr string) (net.Conn, error) {
		return (&net.Dialer{}).DialContext(ctx, "Unix", addr)
	}
	return grpc.DialContext(ctx, addr, grpc.WithInsecure(), grpc.WithContextDialer(d))
}
