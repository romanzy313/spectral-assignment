package client

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/romanzy313/spectral-assignment/pkg/client/sensor"
	"github.com/romanzy313/spectral-assignment/pkg/logger"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

// TODO: use the same logger as server
func Run(ctx context.Context, config Config) {
	e := echo.New()
	e.Logger = logger.New(config.IsDev)

	e.Use(middleware.RequestLogger())
	e.Use(middleware.CORS(config.FrontendOrigin))

	e.GET("/health", func(c *echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// define dependencies here
	sensorClient, err := sensor.NewGrpcClient(config.GrpcServerAddress)
	if err != nil {
		log.Fatalf("failed to initialize client: %s", err.Error())
		return
	}
	defer sensorClient.Close()

	sensorRouter := sensor.NewRouter(sensorClient)
	sensorRouter.Register(e)

	sc := echo.StartConfig{
		Address:         fmt.Sprintf("0.0.0.0:%d", config.Port),
		HideBanner:      true,
		GracefulTimeout: 10 * time.Second, // same as default
	}

	if err := sc.Start(ctx, e); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}

}
