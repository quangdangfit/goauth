package grpc_logr

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	"google.golang.org/grpc"
)

// UnaryServerInterceptor returns a new unary server interceptors that adds zap.Logger to the context.
func UnaryServerInterceptor(logger logr.Logger, opts ...Option) grpc.UnaryServerInterceptor {
	o := evaluateServerOpt(opts)
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		startTime := time.Now()

		resp, err := handler(ctx, req)
		if !o.shouldLog(info.FullMethod, err) {
			return resp, err
		}

		logFinalClientLine(ctx, logger, o, startTime, err, "finished client unary call")

		return resp, err
	}
}

// StreamServerInterceptor returns a new streaming server interceptor that adds zap.Logger to the context.
func StreamServerInterceptor(logger logr.Logger, opts ...Option) grpc.StreamServerInterceptor {
	o := evaluateServerOpt(opts)
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		startTime := time.Now()
		err := handler(srv, stream)
		if !o.shouldLog(info.FullMethod, err) {
			return err
		}
		ctx := stream.Context()
		logFinalClientLine(ctx, logger, o, startTime, err, "finished client unary call")

		return err
	}
}
