package client

import (
	"log"
	"net/http"

	"spectral-assignment/pkg/client/sensor"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

// TODO: use the same logger as server
func Run() {
	e := echo.New()
	// e.Use(middleware.RequestLogger())
	e.Use(middleware.CORS("http://localhost:12002"))

	e.GET("/health", func(c *echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// define dependencies here
	sensorClient, err := sensor.NewGrpcClient("localhost:12000")
	if err != nil {
		log.Fatalf("failed to initialize client: %s", err.Error())
		return
	}
	defer sensorClient.Close()

	sensorRouter := sensor.NewRouter(sensorClient)
	sensorRouter.Register(e)

	if err := e.Start(":12001"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
