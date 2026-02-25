package sensor

import (
	"net/http"

	"github.com/labstack/echo/v5"
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

	e.GET("/api/v1/sensor", func(c *echo.Context) error {
		ctx := c.Request().Context()

		var bindReq GetPageRequestDTO

		err := c.Bind(&bindReq)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		if bindReq.Limit < 1 || bindReq.Limit > 10000 {
			return c.String(http.StatusBadRequest, "limit must be between 1 and 10000")
		}

		req := ToProtoGetPageRequest(bindReq)

		resp, err := r.grpc.GetPage(ctx, req)
		if err != nil {
			c.Logger().Error("failed to get page from server", "error", err)
			return c.String(http.StatusInternalServerError, "internal server error")
		}

		return c.JSON(200, FromProtoGetPageResponse(resp))
	})
}
