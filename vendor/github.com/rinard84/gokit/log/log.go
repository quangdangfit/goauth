package log

import (
	"context"
	"strings"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

// Config holding log config
type Config struct {
	Development       bool   `json:"development"`
	Level             string `json:"level"`
	Encoding          string `json:"encoding"`
	zapOptions        []zap.Option
	contextExtractors []ContextExtractor
	preHooks          []PreLogHook
}

func (c *Config) WithZapOptions(zapOptions ...zap.Option) *Config {
	c.zapOptions = zapOptions
	return c
}

func (c *Config) WithContextExtractors(extractors ...ContextExtractor) *Config {
	c.contextExtractors = extractors
	return c
}

func (c *Config) WithPreLogHook(hook ...PreLogHook) *Config {
	c.preHooks = hook
	return c
}

// DefaultConfig of logger, should set default mode to production to avoid mistake
func DefaultConfig() Config {
	return Config{
		Development: false,
		Level:       "error",
		Encoding:    "console",
	}
}

type levelMap map[string]zapcore.Level

func (lm levelMap) get(level string) (zapcore.Level, bool) {
	if lvl, ok := lm[strings.ToLower(level)]; ok {
		return lvl, ok
	}
	return zapcore.InfoLevel, false
}

var zapLevelMap = levelMap{
	"debug":  zap.DebugLevel,
	"info":   zap.InfoLevel,
	"warn":   zap.WarnLevel,
	"error":  zap.ErrorLevel,
	"dpanic": zap.DPanicLevel,
	"panic":  zap.PanicLevel,
	"fatal":  zap.FatalLevel,
}

func (c Config) BuildAndOverrideDefaultLogger() logr.Logger {
	l := c.Build()
	defaultLogger = l
	return l
}

// Build construct zap Logger by config
func (c Config) Build() logr.Logger {
	zapConfig := zap.NewProductionConfig()

	// set default development log
	if c.Development {
		zapConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	if lvl, ok := zapLevelMap.get(c.Level); ok {
		zapConfig.Level = zap.NewAtomicLevelAt(lvl)
	}

	zapConfig.EncoderConfig.TimeKey = "time"
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapConfig.EncoderConfig.EncodeLevel = func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(levelToColor(level) + level.String() + Reset)
	}
	zapConfig.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	zapConfig.OutputPaths = []string{"stdout"}
	zapConfig.ErrorOutputPaths = []string{"stderr"}

	zapConfig.DisableStacktrace = true
	zapConfig.Encoding = "console"

	z, err := zapConfig.Build(c.zapOptions...)
	if err != nil {
		panic("Who watches the watchmen")
	}

	l := zapr.NewLogger(z)
	logger := &templateLogger{
		l:        l.GetSink(),
		z:        z,
		preHooks: c.preHooks,
	}
	l2 := logr.New(logger)
	// return Logger{l, c.contextExtractors}
	return l2
}

func NewNop() logr.Logger {
	log := zapr.NewLogger(zap.NewNop())
	// var x []ContextExtractor
	// return Logger{log, x}
	return log
}

func New(c *Config) logr.Logger {
	return c.Build()
}

// type Logger struct {
// 	logr.Logger
// 	// WithContext(ctx context.Context) Logger
// }

func NewLogger(c Config) logr.Logger {
	return c.Build()
}

// var _ Logger = &Logger{}

// type Logger struct {
// 	l         logr.Logger
// 	extractor []ContextExtractor
// }

type ContextExtractor func(context.Context) map[string]interface{}

func WithTracingContextValues(l logr.Logger, ctx context.Context) logr.Logger {
	span := trace.SpanFromContext(ctx)
	sc := span.SpanContext()
	ctxValues := make([]interface{}, 0)

	if sc.HasTraceID() {
		ctxValues = append(ctxValues, "traceId", sc.TraceID())
	}

	if sc.HasSpanID() {
		ctxValues = append(ctxValues, "spanId", sc.SpanID())
	}

	// for _, extractor := range log.extractor {
	// 	m := extractor(ctx)
	// 	for k, v := range m {
	// 		ctxValues = append(ctxValues, k, v)
	// 	}
	// }

	return l.WithValues(ctxValues...)
	// return Logger{
	// 	l:         log.l.WithValues(ctxValues...),
	// 	extractor: log.extractor,
	// }
}

var defaultLogger logr.Logger

func init() {
	defaultLogger = DefaultConfig().Build()
}

func DefaultLogger() logr.Logger {
	return defaultLogger
}

// func WithGlobalLogger(ll Logger) {
// 	l = ll
// }

func GetLoggerWithContext(ctx context.Context) logr.Logger {
	return WithTracingContextValues(defaultLogger, ctx)
}

func levelToColor(level zapcore.Level) string {
	switch level {
	case zapcore.DebugLevel:
		return Blue
	case zapcore.InfoLevel:
		return Green
	case zapcore.WarnLevel:
		return Yellow
	case zapcore.ErrorLevel, zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		return Red
	default:
		return Reset
	}
}
