package ecode

import (
	"fmt"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Status struct {
	*status.Status
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// A developer-facing error message, which should be in English. Any
	// user-facing error message should be localized and sent in the
	// [google.rpc.Status.details][google.rpc.Status.details] field, or localized by the client.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Err     string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

// Error new status with code and message
func Error(code Code, message string) *Status {
	return &Status{status.New(codes.Code(code), message), int32(code), message, message}
}

// Errorf new status with code and message
func Errorf(code Code, format string, args ...interface{}) *Status {
	return &Status{status.Newf(codes.Code(code), format, args...), int32(code), fmt.Sprintf(format, args...), fmt.Sprintf(format, args...)}
}

// GRPCStatus()
func (s Status) GRPCStatus() *status.Status {
	return s.Status
}

// Error()
func (s Status) Error() string {
	if m := s.Status.Message(); m != "" {
		return m
	}
	return strconv.Itoa(int(s.Status.Code()))
}
