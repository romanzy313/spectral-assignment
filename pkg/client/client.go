package client

import (
	"context"
	"fmt"
	"net/http"
	"time"

	pb "spectral-assignment/gen/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func getAll(ctx context.Context) (*pb.GetPageResponse, error) {
	conn, err := grpc.NewClient("localhost:12000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %w", err)
	}
	defer conn.Close()

	client := pb.NewSensorServiceClient(conn)

	resp, err := client.GetPage(ctx, &pb.GetPageRequest{
		Cursor: timestamppb.New(time.Time{}),
		Limit:  99999,
	})
	if err != nil {
		return nil, fmt.Errorf("RPC failed: %w", err)
	}

	return resp, nil
}

// TODO: use the same logger as server
func Run() {
	e := echo.New()
	e.Use(middleware.RequestLogger())

	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/all", func(c *echo.Context) error {
		ctx := context.Background()
		_, err := getAll(ctx)
		if err != nil {
			c.Logger().Error("failed to getAll", "error", err.Error())
			return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
		}
		// TODO: wrap into a DTO
		return c.String(500, "Not implemented yet")
	})

	if err := e.Start(":12001"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
