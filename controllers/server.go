package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/quangdangfit/goauth/config"
	api "github.com/quangdangfit/goauth/proto/api/auth"
	repositories "github.com/quangdangfit/goauth/repositories/mysql"
	"github.com/quangdangfit/gokit/library/cache"
	"github.com/quangdangfit/gokit/library/grpc/health"
	"github.com/quangdangfit/gokit/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Controller struct {
	logger   logr.Logger
	cfg      *config.Config
	cache    *cache.Cache
	userRepo repositories.UserRepository

	health.UnimplementedHealthCheckServiceServer
	api.UnimplementedAuthServiceServer
}

func NewController(
	logger logr.Logger,
	cfg *config.Config,
	cache *cache.Cache,
	userRepo repositories.UserRepository,
) *Controller {
	return &Controller{
		logger:   logger,
		cfg:      cfg,
		cache:    cache,
		userRepo: userRepo,
	}
}

// Close implementing service server interface
func (s *Controller) Close(_ context.Context) {}

// Ping implementing service server interface
func (s *Controller) Ping(_ context.Context) error {
	return nil
}

// RegisterWithServer implementing service server interface
func (s *Controller) RegisterWithServer(server *grpc.Server) {
	health.RegisterHealthCheckServiceServer(server, s)
	api.RegisterAuthServiceServer(server, s)
}

// RegisterWithHandler implementing service server interface
func (s *Controller) RegisterWithHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	err := api.RegisterAuthServiceHandler(ctx, mux, conn)
	if err != nil {
		s.logger.Error(err, "errorCount register servers")
	}
	err = health.RegisterHealthCheckServiceHandler(ctx, mux, conn)
	if err != nil {
		s.logger.Error(err, "errorCount register health servers")
	}

	return err
}

// Liveness handle socket is open or not
func (s *Controller) Liveness(ctx context.Context, req *health.LivenessRequest) (*health.LivenessResponse, error) {
	err := s.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return &health.LivenessResponse{
		Content: "ok",
	}, nil
}

// Readiness handle application is ready or not
// this should take into account, saturation of the pod, instead
func (s *Controller) Readiness(ctx context.Context, req *health.ReadinessRequest) (*health.ReadinessResponse, error) {
	if server.IsServerShuttingDown() {
		return nil, status.Error(codes.Unavailable, "service is shutting down")
	}

	return &health.ReadinessResponse{
		Content: "ok",
	}, nil
}
