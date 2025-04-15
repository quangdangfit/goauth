package grpc_ctx

import (
	"context"
	"google.golang.org/grpc"
	"path"
)

// UnaryClientInterceptor returns a new unary client interceptor that optionally logs the execution of external gRPC calls.
func UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		return invoker(newClientCtx(ctx, method), method, req, reply, cc, opts...)
	}
}

// StreamClientInterceptor returns a new streaming client interceptor that optionally logs the execution of external gRPC calls.
func StreamClientInterceptor() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		return streamer(newClientCtx(ctx, method), desc, cc, method, opts...)
	}
}

func newClientCtx(ctx context.Context, fullMethodString string) context.Context {
	service := path.Dir(fullMethodString)[1:]
	method := path.Base(fullMethodString)
	ctx = context.WithValue(ctx, serviceKey, "grpc")
	ctx = context.WithValue(ctx, spanKindKey, "client")
	ctx = context.WithValue(ctx, grpcMethodKey, method)
	ctx = context.WithValue(ctx, grpcServiceKey, service)
	return ctx
}

var clientExtractionSet = []ctxKey{
	serviceKey,
	spanKindKey,
	grpcMethodKey,
	grpcServiceKey,
}

func ExtractClientCtx(ctx context.Context) map[string]interface{} {
	res := make(map[string]interface{})

	for _, k := range clientExtractionSet {
		if v := ctx.Value(k); v != nil {
			res[string(k)] = v
		}
	}

	return res
}
