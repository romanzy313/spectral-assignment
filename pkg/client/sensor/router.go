package sensor

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/romanzy313/spectral-assignment/pkg/common"
)

type Router struct {
	grpc GrpcClient
}

func NewRouter(grpc GrpcClient) *Router {
	return &Router{
		grpc: grpc,
	}
}

func (r *Router) Register(e *echo.Echo) {
	e.GET("/api/v1/sensor/data", r.getSensorData)
}

func (r *Router) getSensorData(c *echo.Context) error {
	ctx := c.Request().Context()

	var bindReq GetPageRequestDTO

	err := c.Bind(&bindReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ApiError{
			Message: "bad request",
		})
	}
	if bindReq.Limit < 1 || bindReq.Limit > 10000 {
		return c.JSON(http.StatusBadRequest, common.ApiError{
			Message: "limit must be between 1 and 10000",
		})
	}

	req := ToProtoGetPageRequest(bindReq)

	resp, err := r.grpc.GetPage(ctx, req)
	if err != nil {
		c.Logger().Error("failed to get page from server", "error", err)
		return c.JSON(http.StatusInternalServerError, common.ApiError{
			Message: "internal server error",
		})
	}

	return c.JSON(200, FromProtoGetPageResponse(resp))
}
