package tracer

import (
	"context"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func UnaryClientInterceptorExcludedHealth(opts ...otelgrpc.Option) grpc.UnaryClientInterceptor {
	interceptor := otelgrpc.UnaryClientInterceptor(opts...)
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		callOpts ...grpc.CallOption,
	) error {
		if method == "/kit.health.HealthCheckService/Readiness" || method == "/kit.health.HealthCheckService/Liveness" {
			return invoker(ctx, method, req, reply, cc, callOpts...)
		} else {
			return interceptor(ctx, method, req, reply, cc, invoker, callOpts...)
		}
	}
}

func UnaryServerInterceptorExcludedHealth(pts ...otelgrpc.Option) grpc.UnaryServerInterceptor {
	interceptor := otelgrpc.UnaryServerInterceptor(pts...)
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		if info != nil {
			if info.FullMethod == "/kit.health.HealthCheckService/Readiness" || info.FullMethod == "/kit.health.HealthCheckService/Liveness" {
				return handler(ctx, req)
			}
		}
		return interceptor(ctx, req, info, handler)
	}
}
