package grpc_logr

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	"github.com/rinard84/gokit/log"
)

func logFinalClientLine(ctx context.Context, logger logr.Logger, o *options, startTime time.Time, err error, msg string) {
	code := o.codeFunc(err)

	logr := log.WithTracingContextValues(logger, ctx).WithValues(
		"grpc.code", code.String(),
		"grpc.time_ms", durationToMilliseconds(time.Since(startTime)),
	)
	if err != nil {
		logr.Error(err, msg)
	} else {
		logr.Info(msg)
	}

}
