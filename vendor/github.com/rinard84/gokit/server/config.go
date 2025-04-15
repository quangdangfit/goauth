package server

import (
	"fmt"
	"net"
)

type ContextKey string

const (
	Identity       ContextKey = "identity"
	StandardClaims ContextKey = "standardclaims"
	// ForbiddenMesage ...
	ForbiddenMesage = "User don't have role to access this api"
)

func (l Listen) String() string {
	return fmt.Sprintf("%s:%d", l.Host, l.Port)
}

// Address represents a network end point address.
type Listen struct {
	Host string `json:"host" mapstructure:"host" yaml:"host"`
	Port int    `json:"port" mapstructure:"port" yaml:"port"`
}

func (a *Listen) CreateListener() (net.Listener, error) {
	lis, err := net.Listen("tcp", a.String())
	if err != nil {
		return nil, fmt.Errorf("failed to listen %s: %w", a.String(), err)
	}
	return lis, nil
}

func createDefaultConfig() *Config {
	config := &Config{
		Grpc:    createDefaultGrpcConfig(),
		Gateway: createDefaultGatewayConfig(),
	}

	return config
}

type Config struct {
	Gateway              *gatewayConfig
	Grpc                 *grpcConfig
	ServiceServers       []ServiceServer
	ShuttingDownCallback func()
}
