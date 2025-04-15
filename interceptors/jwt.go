package interceptors

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	"github.com/quangdangfit/goauth/config"
	"google.golang.org/grpc"
)

type JWT struct {
	cfg    *config.Config
	logger logr.Logger
}

func NewJWT(cfg *config.Config, logger logr.Logger) *JWT {
	return &JWT{
		cfg:    cfg,
		logger: logger,
	}
}

func (i *JWT) Interceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	startTime := time.Now()

	res, err := handler(ctx, req)
	if err != nil {
		i.logger.Error(err, info.FullMethod, "time_ms", time.Since(startTime).Milliseconds())
		return res, err
	}
	i.logger.Info(info.FullMethod, "time_ms", time.Since(startTime).Milliseconds())
	return res, err
}

func (i *JWT) validateToken(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	startTime := time.Now()

	res, err := handler(ctx, req)
	if err != nil {
		i.logger.Error(err, info.FullMethod, "time_ms", time.Since(startTime).Milliseconds())
		return res, err
	}
	i.logger.Info(info.FullMethod, "time_ms", time.Since(startTime).Milliseconds())
	return res, err
}
