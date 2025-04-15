package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-logr/logr"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/quangdangfit/goauth/config"
	"github.com/quangdangfit/goauth/controllers"
	"github.com/quangdangfit/goauth/interceptors"
	"github.com/quangdangfit/goauth/repositories/mysql"
	"github.com/quangdangfit/goauth/utils"
	"github.com/quangdangfit/gokit/library/cache"
	"github.com/quangdangfit/gokit/server"
	"github.com/quangdangfit/gokit/tracing/tracer"
	"github.com/redis/go-redis/v9"
	"github.com/urfave/cli/v2"
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	if err := run(os.Args); err != nil {
		panic(err)
	}
}

// var app *cli.App
var cfg *config.Config

var logger logr.Logger

func run(args []string) error {
	var err error
	cfg, err = config.Load()
	if err != nil {
		return err
	}

	logger = cfg.Log.Build().WithName(cfg.ServiceName)

	if cfg.Env != config.ProductionEnvironment {
		b, _ := json.MarshalIndent(cfg, "", "\t")
		fmt.Println(string(b))
	}

	app := cli.NewApp()
	app.Name = "service"
	app.Commands = []*cli.Command{
		{
			Name:   "server",
			Usage:  "start grpc/http server",
			Action: serverAction,
		},
	}

	if err := app.Run(args); err != nil {
		panic(err)
	}
	return nil
}

func serverAction(_ *cli.Context) error {
	if cfg.Tracing.Jaeger.Enable {
		flush, err := tracer.New(tracer.WithJaegerExporter(tracer.JaegerConfig{
			BaseExporter:      tracer.BaseExporter{Enable: true},
			ServiceName:       cfg.ServiceName,
			CollectorEndpoint: cfg.Tracing.Jaeger.CollectorEndpoint,
			SamplingRate:      cfg.Tracing.Jaeger.SamplingRate,
			Tags:              map[string]string{},
		}))
		if err != nil {
			logger.Error(err, "cannot setup tracer")
		} else {
			defer flush()
		}
	}

	serviceServer, err := InitializeServer(logger, cfg, &cfg.Redis, cfg.ServiceName)
	if err != nil {
		logger.Error(err, "errorCount init servers")
		return err
	}

	s, err := server.New(
		server.WithGrpcAddrListen(cfg.Server.GRPC),
		server.WithGrpcServerUnaryInterceptors(
			interceptors.NewJWT(cfg, logger).Interceptor,
		),
		server.WithGatewayMuxOptions(
			runtime.WithErrorHandler(utils.CustomGrpcError),
			runtime.WithMarshalerOption(
				runtime.MIMEWildcard,
				&runtime.JSONPb{
					MarshalOptions: protojson.MarshalOptions{
						UseProtoNames:   true,
						UseEnumNumbers:  true,
						EmitUnpopulated: true,
					},
					UnmarshalOptions: protojson.UnmarshalOptions{
						DiscardUnknown: true,
					},
				},
			)),
		server.WithServiceServer(serviceServer),
		server.WithGatewayAddrListen(cfg.Server.HTTP),
		server.WithGatewayServerMiddlewares(),
		server.WithGatewayServerHandler(func(muxInput *http.ServeMux) {}),
	)
	if err != nil {
		logger.Error(err, "errorCount new server", err)
		return err
	}

	if err = s.Serve(); err != nil {
		logger.Error(err, "errorCount start server")
		return err
	}
	return nil
}

func InitializeServer(logger logr.Logger, cfg *config.Config, cfgRedis *config.Redis, serviceName string) (*controllers.Controller, error) {
	client, err := NewRedisClient(cfgRedis, logger)
	if err != nil {
		return nil, err
	}
	cacheClient := NewCacheClient(logger, client, serviceName)
	userRepository := mysql.NewUserRepository(logger, cfg, nil)
	return controllers.NewController(logger, cfg, cacheClient, userRepository), nil
}

// wire.go:

func NewRedisClient(cfg *config.Redis, logger logr.Logger) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		DB:       cfg.DB,
		Password: cfg.Password,
	})
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		logger.Error(err, "errorCount pinging Redis")
		return nil, err
	}
	return redisClient, nil
}

func NewCacheClient(logger logr.Logger, redisClient *redis.Client, serviceName string) *cache.Cache {
	return cache.NewCache(logger, redisClient, cache.WithKeyPattern(serviceName+":%s"))
}
