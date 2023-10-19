package app

import (
	"context"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var limiter = rate.NewLimiter(10, 1)

type validatable interface {
	Validate() error
}

func GrpcInterceptor() grpc.ServerOption {
	grpcServerOptions := grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if v, ok := req.(validatable); ok {
			err := v.Validate()
			if err != nil {
				return nil, err
			}
		}
		resp, err = handler(ctx, req)
		return handler(ctx, req)
	})
	return grpcServerOptions
}

func HttpInterceptor() runtime.ServeMuxOption {

	return runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
		meta := map[string]string{}
		if req.Header.Get("Authorization") != "" {
			meta["Authorization"] = req.Header.Get("Authorization")
		}
		if req.Header.Get("x-api-key") != "" {
			meta["x-api-key"] = req.Header.Get("x-api-key")
		}
		if len(meta) == 0 {
			return nil
		}
		return metadata.New(meta)
	})
}

func (m *MicroserviceServer) getUserIdFromToken(ctx context.Context) (string, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	token := md.Get("Authorization")

	if token == nil {
		return "", status.Errorf(codes.PermissionDenied, "user isn't authorized")
	}
	removeBearerFromToken := strings.Split(token[0], " ")
	return removeBearerFromToken[1], nil
}
