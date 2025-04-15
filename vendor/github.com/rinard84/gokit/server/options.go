package server

import (
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// Option configures a gRPC and a gateway server.
type Option func(*Config)

func createConfig(opts []Option) *Config {
	c := createDefaultConfig()
	for _, f := range opts {
		f(c)
	}
	return c
}

func WithOnShuttingDownCallback(callback func()) Option {
	return func(c *Config) {
		c.ShuttingDownCallback = callback
	}
}

// WithGatewayAddr
func WithGatewayAddr(host string, port int) Option {
	return func(c *Config) {
		c.Gateway.Addr = Listen{
			Host: host,
			Port: port,
		}
	}
}

// WithGrpcAddrListen
func WithGatewayAddrListen(l Listen) Option {
	return func(c *Config) {
		c.Gateway.Addr = l
	}
}

// WithGatewayMuxOptions returns an Option that sets runtime.ServeMuxOption(s) to a gateway server.
func WithGatewayMuxOptions(opts ...runtime.ServeMuxOption) Option {
	return func(c *Config) {
		c.Gateway.MuxOptions = append(c.Gateway.MuxOptions, opts...)
	}
}

// WithGatewayServerMiddlewares returns an Option that sets middleware(s) for http.Server to a gateway server.
func WithGatewayServerMiddlewares(middlewares ...HTTPServerMiddleware) Option {
	return func(c *Config) {
		c.Gateway.ServerMiddlewares = append(c.Gateway.ServerMiddlewares, middlewares...)
	}
}

// WithGatewayServerHandler returns an Option that sets hanlers(s) for http.Server to a gateway server.
func WithGatewayServerHandler(handlers ...HTTPServerHandler) Option {
	return func(c *Config) {
		c.Gateway.ServerHandlers = append(c.Gateway.ServerHandlers, handlers...)
	}
}

// WithGatewayServerConfig returns an Option that specifies http.Server configuration to a gateway server.
func WithGatewayServerConfig(cfg *HTTPServerConfig) Option {
	return func(c *Config) {
		c.Gateway.ServerConfig = cfg
	}
}

// WithPassedHeader returns an Option that sets configurations about passed headers for a gateway server.
func WithPassedHeader(decider PassedHeaderDeciderFunc) Option {
	return WithGatewayServerMiddlewares(createPassingHeaderMiddleware(decider))
}

///-------------------------- GRPC options below--------------------------

// WithGrpcAddr
func WithGrpcAddr(host string, port int) Option {
	return func(c *Config) {
		c.Grpc.Addr = Listen{
			Host: host,
			Port: port,
		}
	}
}

// WithGrpcAddrListen
func WithGrpcAddrListen(l Listen) Option {
	return func(c *Config) {
		c.Grpc.Addr = l
	}
}

// WithGrpcServerUnaryInterceptors returns an Option that sets unary interceptor(s) for a gRPC server.
func WithGrpcServerUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) Option {
	return func(c *Config) {
		c.Grpc.ServerUnaryInterceptors = append(c.Grpc.ServerUnaryInterceptors, interceptors...)
	}
}

// WithGrpcServerStreamInterceptors returns an Option that sets stream interceptor(s) for a gRPC server.
func WithGrpcServerStreamInterceptors(interceptors ...grpc.StreamServerInterceptor) Option {
	return func(c *Config) {
		c.Grpc.ServerStreamInterceptors = append(c.Grpc.ServerStreamInterceptors, interceptors...)
	}
}

// WithDefaultLogger returns an Option that sets default grpclogger.LoggerV2 object.
func WithDefaultLogger() Option {
	return func(c *Config) {
		grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stderr))
	}
}

// WithServiceServer
func WithServiceServer(srv ...ServiceServer) Option {
	return func(c *Config) {
		c.ServiceServers = append(c.ServiceServers, srv...)
	}
}

// WithGrpcServerOption
func WithGrpcServerOption(serverOptions ...grpc.ServerOption) Option {
	return func(c *Config) {
		c.Grpc.ServerOption = append(c.Grpc.ServerOption, serverOptions...)
	}
}
