package grpc_ctx

type ctxKey string

const (
	serviceKey             = ctxKey("service")
	spanKindKey            = ctxKey("span.kind")
	grpcServiceKey         = ctxKey("grpc.service")
	grpcMethodKey          = ctxKey("grpc.method")
	grpcStartTimeKey       = ctxKey("grpc.start_time")
	grpcRequestDeadlineKey = ctxKey("grpc.request_deadline")
)
