package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/rinard84/gokit/library/grpc/gatewayopt"

	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"google.golang.org/grpc"
)

// gatewayServer wraps gRPC gateway server setup process.
type gatewayServer struct {
	// Mux *runtime.ServeMux
	// listener *net.Listener
	// mux    *http.Handler
	server *http.Server
	config *gatewayConfig
}

type HTTPServerConfig struct {
	TLSConfig         *tls.Config
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	MaxHeaderBytes    int
	TLSNextProto      map[string]func(*http.Server, *tls.Conn, http.Handler)
	ConnState         func(net.Conn, http.ConnState)
}

func (c *HTTPServerConfig) applyTo(s *http.Server) {
	s.TLSConfig = c.TLSConfig
	s.ReadTimeout = c.ReadTimeout
	s.ReadHeaderTimeout = c.ReadHeaderTimeout
	s.WriteTimeout = c.WriteTimeout
	s.IdleTimeout = c.IdleTimeout
	s.MaxHeaderBytes = c.MaxHeaderBytes
	s.TLSNextProto = c.TLSNextProto
	s.ConnState = c.ConnState
}

type gatewayConfig struct {
	Addr              Listen
	MuxOptions        []runtime.ServeMuxOption
	ServerConfig      *HTTPServerConfig
	ServerMiddlewares []HTTPServerMiddleware
	ServerHandlers    []HTTPServerHandler
}

func createDefaultGatewayConfig() *gatewayConfig {
	config := &gatewayConfig{
		Addr: Listen{
			Host: "0.0.0.0",
			Port: 10080,
		},
		MuxOptions: []runtime.ServeMuxOption{
			gatewayopt.DefaultMarshaler(),
			// runtime.WithProtoErrorHandler(runtime.DefaultHTTPErrorHandler),
		},
		ServerHandlers: []HTTPServerHandler{
			PrometheusHandler,
			PprofHandler,
		},
	}

	return config
}

func newGatewayServer(c *gatewayConfig, conn *grpc.ClientConn, servers []ServiceServer) (*gatewayServer, error) {
	mux := runtime.NewServeMux(c.MuxOptions...)

	var handler http.Handler = mux

	healthFilter := otelhttp.WithFilter(func(r *http.Request) bool {
		return !strings.HasPrefix(r.URL.Path, "/health")
	})
	spanNameFormatter := otelhttp.WithSpanNameFormatter(func(_ string, r *http.Request) string {
		return r.URL.Path
	})
	handler = otelhttp.NewHandler(handler, "otelHttpHandler", healthFilter, spanNameFormatter)

	for i := len(c.ServerMiddlewares) - 1; i >= 0; i-- {
		handler = c.ServerMiddlewares[i](handler)
	}
	httpMux := http.NewServeMux()
	httpMux.Handle("/", handler)
	for _, h := range c.ServerHandlers {
		h(httpMux)
	}

	svr := &http.Server{
		Addr:    c.Addr.String(),
		Handler: httpMux,
	}
	if cfg := c.ServerConfig; cfg != nil {
		cfg.applyTo(svr)
	}

	for _, svr := range servers {
		err := svr.RegisterWithHandler(context.Background(), mux, conn)
		if err != nil {
			return nil, fmt.Errorf("failed to register handler. %w", err)
		}
	}

	return &gatewayServer{
		server: svr,
		config: c,
	}, nil
}

// Serve
func (s *gatewayServer) Serve() error {
	log.Println("http server starting at", s.config.Addr.String())
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Println("Error starting http server, ", err)
		return err
	}

	return nil
}

func (s *gatewayServer) Shutdown(ctx context.Context) {
	// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	// defer cancel()
	err := s.server.Shutdown(ctx)
	log.Println("All http(s) requests finished")
	if err != nil {
		log.Println("failed to shutdown grpc-gateway server: ", err)
	}
}
