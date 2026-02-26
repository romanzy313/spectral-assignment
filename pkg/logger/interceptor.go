package logger

import (
	"context"
	"log/slog"

	"google.golang.org/grpc"
)

func GrpcServerInterceptor(log *slog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		log := log.With(
			"method", info.FullMethod,
		)
		ctx = WithContext(ctx, log)
		return handler(ctx, req)
	}
}
