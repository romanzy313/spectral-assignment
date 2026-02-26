package repository

import (
	"testing"

	"github.com/romanzy313/spectral-assignment/pkg/server/model"
	"github.com/stretchr/testify/assert"
)

func ptr[T any](v T) *T { return &v }

func TestSensorMemoryRepository(t *testing.T) {

	// table tests
	var tests = []struct {
		name           string
		data           []model.SensorData
		cursor         int64
		limit          int32
		expectedResult *model.SensorPage
	}{
		{
			name:           "Empty",
			data:           []model.SensorData{},
			cursor:         0,
			limit:          2,
			expectedResult: &model.SensorPage{Data: []model.SensorData{}, Cursor: nil}},

		{
			name: "From start",
			data: []model.SensorData{
				{Timestamp: 10, Value: 1},
				{Timestamp: 20, Value: 2},
				{Timestamp: 30, Value: 3},
			},
			cursor: 0,
			limit:  2,
			expectedResult: &model.SensorPage{Data: []model.SensorData{
				{Timestamp: 10, Value: 1},
				{Timestamp: 20, Value: 2},
			}, Cursor: ptr(int64(30))}},
		{
			name: "From middle",
			data: []model.SensorData{
				{Timestamp: 10, Value: 1},
				{Timestamp: 20, Value: 2},
				{Timestamp: 30, Value: 3},
				{Timestamp: 40, Value: 4},
			},
			cursor: 20,
			limit:  2,
			expectedResult: &model.SensorPage{Data: []model.SensorData{
				{Timestamp: 20, Value: 2},
				{Timestamp: 30, Value: 3},
			}, Cursor: ptr(int64(40))}},
		{
			name: "Cursor end",
			data: []model.SensorData{
				{Timestamp: 10, Value: 1},
				{Timestamp: 20, Value: 2},
			},
			cursor: 20,
			limit:  1,
			expectedResult: &model.SensorPage{Data: []model.SensorData{
				{Timestamp: 20, Value: 2},
			}, Cursor: nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewSensorMemoryRepository(tt.data)

			got, err := repo.GetPage(t.Context(), tt.cursor, tt.limit)
			assert.NoError(t, err)

			assert.NotNil(t, got)
			assert.Equal(t, tt.expectedResult.Data, got.Data)
			assert.Equal(t, tt.expectedResult.Cursor, got.Cursor)
		})
	}
}
