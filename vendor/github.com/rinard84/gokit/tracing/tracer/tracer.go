package tracer

import (
	"context"

	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

type BaseExporter struct {
	Enable bool `json:"enable" mapstructure:"enable"`
}

type StdoutConfig struct {
	BaseExporter `mapstructure:",squash"`
}

type Exporters struct {
	Jaeger  JaegerConfig          `json:"jaeger" mapstructure:"jaeger"`
	Stdout  StdoutConfig          `json:"stdout" mapstructure:"stdout"`
	GlCloud GlCloudExporterConfig `json:"glcloud" mapstructure:"glcloud"`
}

type TraceConfig struct {
	Exporters `json:"exporters" mapstructure:"exporters"`
}

func DefaultTraceConfig() TraceConfig {
	return TraceConfig{}
}

func (tc TraceConfig) Build() (func(), error) {
	pc := make([]Option, 0)

	if tc.Exporters.Jaeger.Enable {
		pc = append(pc, WithJaegerExporter(tc.Exporters.Jaeger))
	}
	if tc.Exporters.Stdout.Enable {
		pc = append(pc, WithStdoutExporter(stdouttrace.WithPrettyPrint()))
	}
	if tc.Exporters.GlCloud.Enable {
		pc = append(pc, WithGLoudExporter(tc.Exporters.GlCloud))
	}

	return New(pc...)
}

type Config struct {
	providerOptions []sdktrace.TracerProviderOption
	flushes         []func()
}

type Option func(*Config)

func WithProviderOption(providerOption sdktrace.TracerProviderOption) Option {
	return func(config *Config) {
		config.providerOptions = append(config.providerOptions, providerOption)
	}
}

func WithStdoutExporter(opts ...stdouttrace.Option) Option {
	// Create stdout exporter to be able to retrieve the collected spans.

	exporter, err := stdouttrace.New(opts...)
	if err != nil {
		return func(config *Config) {}
	}
	return func(config *Config) {
		config.providerOptions = append(config.providerOptions, sdktrace.WithSyncer(exporter))
	}
}

type GlCloudExporterConfig struct {
	BaseExporter `mapstructure:",squash"`
	ProjectId    string            `json:"project_id" mapstructure:"project_id"`
	ServiceName  string            `json:"service_name" mapstructure:"service_name"`
	SamplingRate *float64          `json:"sampling_rate" mapstructure:"sampling_rate"`
	Tags         map[string]string `json:"tags" mapstructure:"tags"`
}

func WithGLoudExporter(cfg GlCloudExporterConfig) Option {
	return func(config *Config) {
		exporter, err := texporter.New(texporter.WithProjectID(cfg.ProjectId))
		if err != nil {
			return
		}
		config.flushes = append(config.flushes, func() {
			_ = exporter.Shutdown(context.Background()) //nolint:errcheck
		})
		kv := make([]attribute.KeyValue, 0, len(cfg.Tags)+1)
		kv = append(kv, semconv.ServiceNameKey.String(cfg.ServiceName))
		for k, v := range cfg.Tags {
			kv = append(kv, attribute.KeyValue{
				Key:   attribute.Key(k),
				Value: attribute.StringValue(v),
			})
		}
		rate := 0.1
		if cfg.SamplingRate != nil {
			rate = *cfg.SamplingRate
		}
		config.providerOptions = append(config.providerOptions,
			sdktrace.WithBatcher(exporter),
			sdktrace.WithResource(resource.NewSchemaless(kv...)),
			sdktrace.WithSampler(sdktrace.TraceIDRatioBased(rate)),
		)
	}
}

type JaegerConfig struct {
	BaseExporter      `mapstructure:",squash"`
	CollectorEndpoint string            `json:"collector_endpoint" mapstructure:"collector_endpoint"`
	ServiceName       string            `json:"service_name" mapstructure:"service_name"`
	SamplingRate      *float64          `json:"sampling_rate" mapstructure:"sampling_rate"`
	Tags              map[string]string `json:"tags" mapstructure:"tags"`
}

func WithJaegerExporter(cfg JaegerConfig) Option {
	return func(config *Config) {
		kv := make([]attribute.KeyValue, 1, len(cfg.Tags)+1)
		kv = append(kv, semconv.ServiceNameKey.String(cfg.ServiceName))
		for k, v := range cfg.Tags {
			kv = append(kv, attribute.KeyValue{
				Key:   attribute.Key(k),
				Value: attribute.StringValue(v),
			})
		}

		opts := make([]jaeger.CollectorEndpointOption, 0)
		if cfg.CollectorEndpoint != "" {
			opts = append(opts, jaeger.WithEndpoint(cfg.CollectorEndpoint))
		}
		exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(opts...))
		if err != nil {
			return
		}
		config.flushes = append(config.flushes, func() {
			_ = exporter.Shutdown(context.Background()) //nolint:errcheck
		})
		rate := 0.1
		if cfg.SamplingRate != nil {
			rate = *cfg.SamplingRate
		}
		config.providerOptions = append(config.providerOptions,
			sdktrace.WithSyncer(exporter),
			sdktrace.WithResource(resource.NewSchemaless(kv...)),
			sdktrace.WithSampler(sdktrace.TraceIDRatioBased(rate)),
		)
	}
}

func defaultProviderConfig() []sdktrace.TracerProviderOption {
	opts := make([]sdktrace.TracerProviderOption, 0, 1)
	opts = append(opts, sdktrace.WithSampler(sdktrace.AlwaysSample()))

	return opts
}

func defaultConfig() *Config {
	return &Config{
		providerOptions: defaultProviderConfig(),
	}
}

//Flush flush leftover data
func (cfg *Config) Flush() {
	for _, fn := range cfg.flushes {
		fn()
	}
}

//New register new global trace provider with options
func New(opts ...Option) (func(), error) {
	cfg := defaultConfig()
	for _, f := range opts {
		f(cfg)
	}

	// For the demonstration, use sdktrace.AlwaysSample sampler to sample all traces.
	// In a production application, use sdktrace.ProbabilitySampler with a desired probability.
	tp := sdktrace.NewTracerProvider(
		cfg.providerOptions...,
	)

	otel.SetTracerProvider(tp)
	return func() {
		cfg.Flush()
	}, nil
}

// StartSpan wrap func global.TraceProvider().Tracer() of library 'otel/api/trace',
// it will be help easy and clean to use when caller call tracing with global tracer was existed
func StartSpan(ctx context.Context, spanName string) (context.Context, trace.Span) {
	return otel.Tracer("").Start(ctx, spanName)
}
