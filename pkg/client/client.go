package client

import (
	"fmt"
	"log"
	"net/http"

	"github.com/romanzy313/spectral-assignment/pkg/client/sensor"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

// TODO: use the same logger as server
func Run(config Config) {
	e := echo.New()
	// e.Use(middleware.RequestLogger())
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

	addr := fmt.Sprintf("0.0.0.0:%d", config.Port)
	if err := e.Start(addr); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}

}
