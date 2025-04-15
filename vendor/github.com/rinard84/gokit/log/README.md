# Log

## With trace context

Init the logger
```
	cfg := log.DefaultConfig()
	l := cfg.MustBuildLogR()
	lr := log.NewLogger(l)
```

Add more context for a good logging

```
	lr = lr.WithContext(ctx)
	lr.Info("some error")
	lr.Error(nil, "some error")
```

 Output
 
 ```
$ {"level":"info","ts":1602476155.1094582,"caller":"examples/main.go:24","msg":"some error","TraceId":"5c0bee791083ebc43e26cdee3102c99b","SpanId":"8c14a340f17cf3b1"}
$ {"level":"error","ts":1602476155.1094832,"caller":"examples/main.go:25","msg":"some error","TraceId":"5c0bee791083ebc43e26cdee3102c99b","SpanId":"8c14a340f17cf3b1","stacktrace":"github.com/go-logr/zapr.(*zapLogger).Error\n\t/home/vchitai/code/gato/vendor/github.com/go-logr/zapr/zapr.go:132\nmain.main\n\t/home/vchitai/code/github.com/rinard84/gokit/library/log/examples/main.go:25\nruntime.main\n\t/usr/local/go/src/runtime/proc.go:203"}
```