package gatewayopt

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/url"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

//https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/customizing_your_gateway/

// DefaultMarshaler return default grpc-gateway marshaler with additional support for emit empty field
func DefaultMarshaler() runtime.ServeMuxOption {
	// return runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true})
	return runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})
}

// ProtoJSONMarshaler return the marshaler option with support serialization data with json_name specific
// func ProtoJSONMarshaler() runtime.ServeMuxOption {
// 	return runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{EmitDefaults: true})
// }

// Redirect support redirect endpoint for GRPC gateway by using server metadata
// Example:
//		redirect := metadata.Pairs(
//			"Redirect", "https://5kbps.io",
//			"Redirect-Code", "308",
//		)
//		grpc.SendHeader(ctx, redirect)
//
// This will redirect endpoint to https://5kbps.io with code 308, by default redirect code is 301.
// func Redirect() runtime.ServeMuxOption {
// 	fn := func(ctx context.Context, w http.ResponseWriter, _ protoiface.MessageV1) error {
// 		md, ok := runtime.ServerMetadataFromContext(ctx)
// 		if !ok {
// 			return nil
// 		}

// 		redirect := md.HeaderMD.Get("redirect")
// 		redirectCode := md.HeaderMD.Get("redirect-code")
// 		if len(redirect) > 0 {
// 			w.Header().Set("Location", redirect[0])
// 			if len(redirectCode) > 0 {
// 				code, err := strconv.Atoi(redirectCode[0])
// 				if err != nil {
// 					return err
// 				}
// 				if code > 300 && code < 400 {
// 					w.WriteHeader(code)
// 				}
// 			}

// 			w.WriteHeader(301)
// 		}
// 		return nil
// 	}
// 	return runtime.WithForwardResponseOption(fn)
// }

// FormURLEncodedMarshaler is custom Marshaler that supports reading x-www-form-urlencoded mine type
func FormURLEncodedMarshaler() runtime.ServeMuxOption {
	return runtime.WithMarshalerOption("application/x-www-form-urlencoded", &formMarshaler{
		&runtime.JSONPb{MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		}},
	})
}

type formMarshaler struct {
	*runtime.JSONPb
}

func (j *formMarshaler) NewDecoder(r io.Reader) runtime.Decoder {
	return runtime.DecoderFunc(func(v interface{}) error { return formDecoderFunc(r, v) })
}

func formDecoderFunc(d io.Reader, v interface{}) error {
	msg, ok := v.(proto.Message)
	if !ok {
		return fmt.Errorf("not proto message")
	}

	formData, err := ioutil.ReadAll(d)
	if err != nil {
		return err
	}
	values, err := url.ParseQuery(string(formData))
	if err != nil {
		return err
	}
	filter := &utilities.DoubleArray{}
	// err = runtime.PopulateQueryParameters(proto.MessageV2(msg), values, filter)
	err = runtime.PopulateQueryParameters(msg, values, filter)
	if err != nil {
		return err
	}
	return nil
}
