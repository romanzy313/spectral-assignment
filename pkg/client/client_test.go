package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v5"
	protov1 "github.com/romanzy313/spectral-assignment/gen/proto/v1"
	"github.com/romanzy313/spectral-assignment/pkg/client/sensor"

	"github.com/stretchr/testify/assert"
)

func ptr[T any](v T) *T { return &v }

type mockTestClient struct {
	sensor.GrpcClient
}

func (m *mockTestClient) Close() error {
	return nil
}

func (m *mockTestClient) GetPage(ctx context.Context, req *protov1.GetPageRequest) (*protov1.GetPageResponse, error) {
	return &protov1.GetPageResponse{NextCursor: ptr(int64(42)), Data: []*protov1.SensorData{
		{Timestamp: 10, Value: 1},
	}}, nil
}

func TestClient(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		method         string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "valid request",
			url:            "/api/v1/sensor/data?cursor=0&limit=2",
			method:         http.MethodGet,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"data":[{"t":10,"v":1}],"nextCursor":42}`,
		},
		{
			name:           "bad limit",
			url:            "/api/v1/sensor/data?cursor=0&limit=99999",
			method:         http.MethodGet,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"message":"limit must be between 1 and 10000"}`,
		},
	}

	e := echo.New()

	mockClient := &mockTestClient{}

	sensorRouter := sensor.NewRouter(mockClient)
	sensorRouter.Register(e)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			rec := httptest.NewRecorder()

			e.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)

			assert.JSONEq(t, tt.expectedBody, rec.Body.String())
		})
	}

}
