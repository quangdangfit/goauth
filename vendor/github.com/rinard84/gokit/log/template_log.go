package log

import (
	"fmt"
	"strings"

	"github.com/go-logr/logr"
	"github.com/iancoleman/strcase"
	"go.uber.org/zap"
)

type PreLogHook func(level string, msg string, keysAndValues ...interface{}) (string, []interface{})

type templateLogger struct {
	l        logr.LogSink
	preHooks []PreLogHook
	z        *zap.Logger
}

func (c *templateLogger) Init(ri logr.RuntimeInfo) {
	c.l.Init(ri)
}

func (c *templateLogger) Enabled(level int) bool {
	return c.l.Enabled(level)
}

func (c *templateLogger) Info(level int, msg string, keysAndValues ...interface{}) {
	if c.Enabled(level) {
		if len(c.preHooks) > 0 {
			for _, v := range c.preHooks {
				msg, keysAndValues = v("info", msg, keysAndValues...)
			}
		}
		c.l.Info(level, msg, keysAndValues...)
	}
}

func (c *templateLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	if checkedEntry := c.z.Check(zap.ErrorLevel, msg); checkedEntry != nil {
		if len(c.preHooks) > 0 {
			for _, v := range c.preHooks {
				msg, keysAndValues = v("error", msg, keysAndValues...)
			}
		}
		c.l.Error(err, msg, keysAndValues...)
	}
}

func (c *templateLogger) WithValues(keysAndValues ...interface{}) logr.LogSink {
	newLogger := *c
	newLogger.l = c.l.WithValues(keysAndValues...)
	return &newLogger
}

func (c *templateLogger) WithName(name string) logr.LogSink {
	newLogger := *c
	newLogger.l = c.l.WithName(name)
	return &newLogger

}

func noOpHook(level string, msg string, keysAndValues ...interface{}) (string, []interface{}) {
	return msg, keysAndValues
}

func CreatePanicCheckCamelCaseKeyByEnv(env string) PreLogHook {
	if env == "prod" {
		return noOpHook
	}

	return func(level string, msg string, keysAndValues ...interface{}) (string, []interface{}) {
		for i := 0; i < len(keysAndValues)/2; i++ {
			key := keysAndValues[i*2]
			if strings.Contains(key.(string), "_") {
				panic(fmt.Sprintf("unsupport snake case key log: %v", key))
			}
		}
		return msg, keysAndValues
	}
}

func ForceCamelCaseKeyHook(level string, msg string, keysAndValues ...interface{}) (string, []interface{}) {
	if len(keysAndValues)%2 != 0 {
		fmt.Println("key value is not pair")
	} else {
		for i := 0; i < len(keysAndValues)/2; i++ {
			key := keysAndValues[i*2]
			camelKey := strcase.ToLowerCamel(key.(string))
			keysAndValues[i*2] = camelKey
		}
	}
	return msg, keysAndValues
}

func CreateWhitelistMaskingSensitiveDataKeysHook(keys ...string) PreLogHook {
	whitelistData := make(map[string]struct{}, len(keys))
	var empty = struct{}{}
	for _, v := range keys {
		whitelistData[v] = empty
	}
	return func(level string, msg string, keysAndValues ...interface{}) (string, []interface{}) {
		if len(keysAndValues)%2 != 0 {
			fmt.Println("key value is not pair")
		} else {
			for i := 0; i < len(keysAndValues)/2; i++ {
				key := keysAndValues[i*2]
				if _, ok := whitelistData[key.(string)]; ok {
					value := keysAndValues[i*2+1]
					value = maskingSensitiveData(value)
					keysAndValues[i*2+1] = value
				}

				camelKey := strcase.ToLowerCamel(key.(string))
				keysAndValues[i*2] = camelKey
			}
		}
		return msg, keysAndValues
	}
}

func maskingSensitiveData(data interface{}) interface{} {
	switch t := data.(type) {
	case string:
		v := []rune(t)
		for i := 2; i < len(v)-2; i++ {
			v[i] = '*'
		}
		return string(v)
	//TODO handle additional data type
	default:
		return data
	}
}
