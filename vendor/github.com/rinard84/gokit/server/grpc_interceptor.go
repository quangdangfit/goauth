package server

import (
	"context"

	"google.golang.org/grpc"
)

func NoCancelInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		return handler(withoutCancel(ctx), req)
	}
}

type noCancel struct {
	context.Context
}

func (c noCancel) Done() <-chan struct{} { return nil }
func (c noCancel) Err() error            { return nil }

// withoutCancel returns a context that is never canceled.
func withoutCancel(ctx context.Context) context.Context {
	return noCancel{Context: ctx}
}
