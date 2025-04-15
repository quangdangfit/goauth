package grpc_ctx

import (
	"context"
	"google.golang.org/grpc"
	"path"
	"time"
)

// UnaryServerInterceptor returns a new unary client interceptor that optionally logs the execution of external gRPC calls.
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		return handler(newServerCtx(ctx, info.FullMethod), req)
	}
}

// StreamServerInterceptor returns a new streaming client interceptor that optionally logs the execution of external gRPC calls.
func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		return handler(srv, ss)
	}
}

func newServerCtx(ctx context.Context, fullMethodString string) context.Context {
	service := path.Dir(fullMethodString)[1:]
	method := path.Base(fullMethodString)
	ctx = context.WithValue(ctx, grpcStartTimeKey, time.Now().Format(time.RFC3339))
	if d, ok := ctx.Deadline(); ok {
		ctx = context.WithValue(ctx, grpcRequestDeadlineKey, d.Format(time.RFC3339))
	}
	ctx = context.WithValue(ctx, serviceKey, "grpc")
	ctx = context.WithValue(ctx, spanKindKey, "server")
	ctx = context.WithValue(ctx, grpcMethodKey, method)
	ctx = context.WithValue(ctx, grpcServiceKey, service)
	return ctx
}

var serverExtractionSet = []ctxKey{
	serviceKey,
	spanKindKey,
	grpcMethodKey,
	grpcServiceKey,
	grpcStartTimeKey,
	grpcRequestDeadlineKey,
}

func ExtractServerCtx(ctx context.Context) map[string]interface{} {
	res := make(map[string]interface{})

	for _, k := range serverExtractionSet {
		if v := ctx.Value(k); v != nil {
			res[string(k)] = v
		}
	}

	return res
}
