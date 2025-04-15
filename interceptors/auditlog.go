package interceptors

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	"github.com/rinard84/auth-service/config"
	"google.golang.org/grpc"
)

type AuditLog struct {
	cfg    *config.Config
	logger logr.Logger
}

func NewAuditLog(cfg *config.Config, logger logr.Logger) *AuditLog {
	return &AuditLog{
		cfg:    cfg,
		logger: logger,
	}
}

func (i *AuditLog) Interceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	startTime := time.Now()
	res, err := handler(ctx, req)
	if err != nil {
		i.logger.Error(err, info.FullMethod, "time_ms", time.Since(startTime).Milliseconds())
		return res, err
	}
	i.logger.Info(info.FullMethod, "time_ms", time.Since(startTime).Milliseconds())
	return res, err
}
