package server

import (
	"context"
	"fmt"
	"log"

	grpcCtxTags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	grpcCtx "github.com/rinard84/gokit/library/grpc/ctx"
	grpcLogr "github.com/rinard84/gokit/library/grpc/logging"
	log2 "github.com/rinard84/gokit/log"
	"github.com/rinard84/gokit/tracing/tracer"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

type grpcConfig struct {
	Addr                     Listen
	ServerUnaryInterceptors  []grpc.UnaryServerInterceptor
	ServerStreamInterceptors []grpc.StreamServerInterceptor
	ServerOption             []grpc.ServerOption
	MaxConcurrentStreams     uint32
	MaxRecvMsgSize           int
	MaxSendMsgSize           int
}

func createDefaultGrpcConfig() *grpcConfig {

	c := log2.DefaultConfig()
	c.WithContextExtractors(grpcCtx.ExtractServerCtx,
		grpcLogr.TagsToFields)
	lgr := log2.NewLogger(c)

	grpcPrometheus.EnableHandlingTimeHistogram()
	config := &grpcConfig{
		Addr: Listen{
			Host: "0.0.0.0",
			Port: 10443,
		},
		ServerUnaryInterceptors: []grpc.UnaryServerInterceptor{
			NoCancelInterceptor(),
			tracer.UnaryServerInterceptorExcludedHealth(),
			grpcPrometheus.UnaryServerInterceptor,
			grpcCtxTags.UnaryServerInterceptor(grpcCtxTags.WithFieldExtractor(grpcCtxTags.CodeGenRequestFieldExtractor)),
			grpcCtx.UnaryServerInterceptor(),
			grpcLogr.UnaryServerInterceptor(lgr),
			//grpc_zap.UnaryServerInterceptor(logger),
			grpcValidator.UnaryServerInterceptor(),
		},
		ServerStreamInterceptors: []grpc.StreamServerInterceptor{
			otelgrpc.StreamServerInterceptor(),
			grpcPrometheus.StreamServerInterceptor,
			grpcCtxTags.StreamServerInterceptor(grpcCtxTags.WithFieldExtractor(grpcCtxTags.CodeGenRequestFieldExtractor)),
			grpcCtx.StreamServerInterceptor(),
			grpcLogr.StreamServerInterceptor(lgr),
			//grpc_zap.StreamServerInterceptor(logger),
			grpcValidator.StreamServerInterceptor(),
		},

		MaxConcurrentStreams: 1000,
		MaxRecvMsgSize:       10 * 1024 * 1024, // 10 MB
		MaxSendMsgSize:       10 * 1024 * 1024, // 10 MB
	}

	return config
}

func (c *grpcConfig) ServerOptions() []grpc.ServerOption {
	return append(
		[]grpc.ServerOption{
			grpc.ChainUnaryInterceptor(c.ServerUnaryInterceptors...),
			grpc.ChainStreamInterceptor(c.ServerStreamInterceptors...),
			// grpc_middleware.WithUnaryServerChain(c.ServerUnaryInterceptors...),
			// grpc_middleware.WithStreamServerChain(c.ServerStreamInterceptors...),
			grpc.MaxConcurrentStreams(c.MaxConcurrentStreams),
			grpc.MaxRecvMsgSize(c.MaxRecvMsgSize),
			grpc.MaxSendMsgSize(c.MaxSendMsgSize),
		},
		c.ServerOption...,
	)
}

// grpcServer wraps grpc.Server setup process.
type grpcServer struct {
	server *grpc.Server
	config *grpcConfig
}

func newGrpcServer(c *grpcConfig, servers []ServiceServer) *grpcServer {
	s := grpc.NewServer(c.ServerOptions()...)
	for _, svr := range servers {
		svr.RegisterWithServer(s)
	}
	grpcPrometheus.Register(s)
	return &grpcServer{
		server: s,
		config: c,
	}
}

// Serve implements Server.Server
func (s *grpcServer) Serve() error {
	l, err := s.config.Addr.CreateListener()
	if err != nil {
		return fmt.Errorf("failed to create listener %w", err)
	}

	log.Println("gRPC server is starting ", l.Addr())

	err = s.server.Serve(l)

	if err != nil {
		log.Println(err)
		return fmt.Errorf("failed to serve gRPC server %w", err)
	}
	log.Println("gRPC server ready")

	return nil
}

// Shutdown
func (s *grpcServer) Shutdown(ctx context.Context) {
	s.server.GracefulStop()
}
