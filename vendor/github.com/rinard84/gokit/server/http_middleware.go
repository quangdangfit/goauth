package server

import (
	"net/http"
	"sync"

	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// HTTPServerMiddleware is an interface of http server middleware
type HTTPServerMiddleware func(http.Handler) http.Handler

// PassedHeaderDeciderFunc returns true if given header should be passed to gRPC server metadata.
type PassedHeaderDeciderFunc func(string) bool

func createPassingHeaderMiddleware(decide PassedHeaderDeciderFunc) HTTPServerMiddleware {
	return func(next http.Handler) http.Handler {
		cache := new(sync.Map)
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			newHeader := make(http.Header, 2*len(r.Header))

			for k := range r.Header {
				v := r.Header.Get(k)
				if newKey, ok := cache.Load(k); ok {
					newHeader.Set(newKey.(string), v)
				} else if decide(k) {
					newKey := runtime.MetadataHeaderPrefix + k
					cache.Store(k, newKey)
					newHeader.Set(newKey, v)
				}
				newHeader.Set(k, v)
			}

			r.Header = newHeader

			next.ServeHTTP(w, r)
		})
	}
}
