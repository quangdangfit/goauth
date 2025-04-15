package health

import (
	"strings"

	"google.golang.org/grpc"
)

func IsHealthCheck(info *grpc.UnaryServerInfo) bool {
	ret := strings.HasPrefix(info.FullMethod, "/kit.health.")
	return ret
}
