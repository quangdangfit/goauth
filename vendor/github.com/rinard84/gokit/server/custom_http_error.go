package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rinard84/gokit/library/ecode"
	"google.golang.org/grpc/status"
)

type ErrorBody struct {
	Msg  string `json:"message,omitempty"`
	Code uint32 `json:"code,omitempty"`
}

func CustomHTTPError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	const fallback = `{"code":0, "message":"Unknow error"}`

	contentType := "application/json"
	errCode := uint32(0)
	errDes := "Unknow error"
	stC := status.Convert(err)
	if stC != nil {
		errCode = uint32(stC.Code())
		errDes = stC.Message()
		// overwrite status code 3
		if errCode == 3 {
			errCode = uint32(ecode.InvalidInput.Code)
			errDes = ecode.InvalidInput.Message
		}
	}
	w.Header().Set("Content-type", contentType)
	jErr := json.NewEncoder(w).Encode(ErrorBody{
		Msg:  errDes,
		Code: errCode,
	})

	if jErr != nil {
		w.Write([]byte(fallback))
	}
}
