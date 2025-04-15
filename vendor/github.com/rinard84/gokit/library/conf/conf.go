package conf

import (
	"github.com/rinard84/gokit/log"
	"github.com/rinard84/gokit/server"
	"github.com/rinard84/gokit/tracing/tracer"
)

// deploy env.
const (
	DeployEnvDev  = "dev"
	DeployEnvQa   = "qa"
	DeployEnvUat  = "uat"
	DeployEnvProd = "prod"
)

// Config hold http/grpc server config
type ServerConfig struct {
	GRPC server.Listen `json:"grpc" mapstructure:"grpc" yaml:"grpc"`
	HTTP server.Listen `json:"http" mapstructure:"http" yaml:"http"`
}

// DefaultServerConfig return a default server config
func DefaultServerConfig() ServerConfig {
	return ServerConfig{
		GRPC: server.Listen{
			Host: "0.0.0.0",
			Port: 40000,
		},
		HTTP: server.Listen{
			Host: "0.0.0.0",
			Port: 4000,
		},
	}
}

type SecretConfig struct {
	Enable     bool   `json:"enable" mapstructure:"enable"`
	FullPath   string `json:"full_path" mapstructure:"full_path"`
	MountPath  string `json:"mount_path" mapstructure:"mount_path"`
	SecretPath string `json:"secret_path" mapstructure:"secret_path"`
}

// Config ...
type Base struct {
	Env string     `json:"env" mapstructure:"env"`
	Log log.Config `json:"log" mapstructure:"log"`
	// LogLevel int `json:"log_level" mapstructure: "log_level"`
	Server       ServerConfig       `json:"server" mapstructure:"server"`
	Tracing      tracer.TraceConfig `json:"tracing" mapstructure:"tracing"`
	SecretConfig SecretConfig       `json:"secret_config" mapstructure:"secret_config"`
}

func DefaultBaseConfig() *Base {
	return &Base{
		Env: DeployEnvDev,
		// LogLevel: 2,
		Log:     log.DefaultConfig(),
		Server:  DefaultServerConfig(),
		Tracing: tracer.DefaultTraceConfig(),
	}
}
