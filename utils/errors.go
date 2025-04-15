package utils

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/status"
)

type ErrorBody struct {
	Success bool   `json:"success" protobuf:"bytes,1,opt,name=success,proto3"`
	Msg     string `json:"message" protobuf:"bytes,1,opt,name=message,proto3"`
	Code    string `json:"code,omitempty" protobuf:"bytes,1,opt,name=code,proto3"`
}

func CustomGrpcError(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	if strings.HasPrefix(r.URL.Path, "/health") {
		runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, w, r, err)
		return
	}
	contentType := "application/json"
	errCode := uint32(0)
	errDes := "Invalid data"
	stC := status.Convert(err)
	httpResponse := http.StatusBadRequest
	bodyRes := ErrorBody{
		Msg:     errDes,
		Success: false,
	}
	if stC != nil {
		errCode = uint32(stC.Code())
		if errCode >= 100 && errCode <= 599 {
			httpResponse = int(errCode)
		}
		bodyRes.Msg = stC.Message()
		switch errDes {
		case "MANY_UNSUCCESSFUL":
			bodyRes.Msg = "Too many unsuccessful code"
			bodyRes.Code = "MANY_UNSUCCESSFUL"
		}
	}
	w.Header().Set("Content-type", contentType)
	w.WriteHeader(httpResponse)
	jErr := json.NewEncoder(w).Encode(bodyRes)

	if jErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"success": false, "message":"Bad Request", "code": "BAD_REQUEST"}`))
		return
	}
}
