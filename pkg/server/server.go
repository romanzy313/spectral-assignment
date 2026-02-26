package server

import (
	"context"
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/romanzy313/spectral-assignment/pkg/logger"
	"github.com/romanzy313/spectral-assignment/pkg/server/repository"
	"github.com/romanzy313/spectral-assignment/pkg/server/router"
	"github.com/romanzy313/spectral-assignment/pkg/server/service"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
)

func Run(ctx context.Context, config Config) {
	log := logger.New(config.IsDev)

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			// TODO: add more observability
			logger.GrpcServerInterceptor(log),
			grpc_recovery.UnaryServerInterceptor(),
		),
	)

	// dependencies are initialized here
	mockData, err := repository.ReadCsvData("./meterusage.csv")
	if err != nil {
		log.Error("failed to read csv data", "error", err)
		os.Exit(1)
	}
	sensorRepo := repository.NewSensorMemoryRepository(mockData)
	sensorService := service.NewSensorService(sensorRepo)
	sensorRouter := router.NewSensorRouter(sensorService)
	sensorRouter.Register(s)

	addr := fmt.Sprintf("0.0.0.0:%d", config.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Error("failed to listen", "error", err)
		os.Exit(1)
	}

	go func() {
		<-ctx.Done()
		log.Debug("starting graceful shutdown")

		s.GracefulStop()
	}()

	log.Info("server listening", "port", config.Port)
	if err := s.Serve(lis); err != nil {
		log.Error("failed to serve", "error", err)
		os.Exit(1)
	}

	log.Debug("server shutdown")
}
