package server

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rinard84/gokit/tracing/tracer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	isShuttingDown = false
)

// Server is the framework instance.
type Server struct {
	grpcServer    *grpcServer
	gatewayServer *gatewayServer
	config        *Config
	globalCtx     context.Context
}

// New creates a server intstance.
func New(opts ...Option) (*Server, error) {
	c := createConfig(opts)

	log.Println("Create grpc server")
	grpcServer := newGrpcServer(c.Grpc, c.ServiceServers)
	// if err != nil {
	// 	return nil, fmt.Errorf("Faild to create grpc server. %w", err)
	// }

	conn, err := grpc.Dial(c.Grpc.Addr.String(), grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024*1024*50)),
		grpc.WithChainUnaryInterceptor(
			tracer.UnaryClientInterceptorExcludedHealth(),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("fail to dial gRPC server. %w", err)
	}

	log.Println("Create gateway server")
	gatewayServer, err := newGatewayServer(c.Gateway, conn, c.ServiceServers)
	if err != nil {
		return nil, fmt.Errorf("fail to create gateway server. %w", err)
	}

	return &Server{
		grpcServer:    grpcServer,
		gatewayServer: gatewayServer,
		config:        c,
	}, nil
}

// Serve starts gRPC and Gateway servers.
func (s *Server) Serve() error {
	stop := make(chan os.Signal, 1)
	errch := make(chan error)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.gatewayServer.Serve(); err != nil {
			log.Println("Error starting http server, ", err)
			errch <- err
		}
	}()

	go func() {
		if err := s.grpcServer.Serve(); err != nil {
			log.Println("Error starting gRPC server, ", err)
			errch <- err
		}
	}()

	// shutdown
	for {
		select {
		case <-stop:
			log.Println("Received sigterm")
			isShuttingDown = true
			if s.config.ShuttingDownCallback != nil {
				s.config.ShuttingDownCallback()
			}
			time.Sleep(20 * time.Second) // wait for K8S remove endpoint
			log.Println("Shutting down server")
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			for _, ss := range s.config.ServiceServers {
				ss.Close(ctx)
			}

			s.gatewayServer.Shutdown(ctx)
			s.grpcServer.Shutdown(ctx)
			return nil
		case err := <-errch:
			return err
		}
	}
}

func IsServerShuttingDown() bool {
	return isShuttingDown
}

func init() {
	job := os.Getenv("JOB")
	istioEnable := os.Getenv("ISTIO_ENABLE")
	if job == "true" && istioEnable == "true" {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		go func(signal chan os.Signal) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("recover from disable Istio", r)
				}
			}()
			select {
			case <-signal:
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()
				cleanUpIstio(ctx)
			}
		}(stop)
	}
}

func cleanUpIstio(ctx context.Context) {
	fmt.Println("Stop istio by posting to http://127.0.0.1:15000/quitquitquit")
	request, _ := http.NewRequestWithContext(ctx, "POST", "http://127.0.0.1:15000/quitquitquit", bytes.NewBufferString("{}"))
	request.Header.Set("Content-Type", "application/json")
	_, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("Fail to send quit signal to istio ", err)
	} else {
		fmt.Println("Successfully send term signal to istio")
	}
}
