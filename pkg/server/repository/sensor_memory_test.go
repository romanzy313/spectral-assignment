package repository

import (
	"reflect"
	"testing"

	"github.com/romanzy313/spectral-assignment/pkg/server/model"
)

func ptr[T any](v T) *T { return &v }

// table tests
var cases = []struct {
	name   string
	data   []model.SensorData
	cursor int64
	limit  int32
	result *model.SensorPage
}{
	{"Empty", []model.SensorData{}, 0, 2, &model.SensorPage{Data: []model.SensorData{}, Cursor: nil}},
	{"From start", []model.SensorData{
		{Timestamp: 10, Value: 1},
		{Timestamp: 20, Value: 2},
		{Timestamp: 30, Value: 3},
	}, 0, 2, &model.SensorPage{Data: []model.SensorData{
		{Timestamp: 10, Value: 1},
		{Timestamp: 20, Value: 2},
	}, Cursor: ptr(int64(30))}},
	{"From middle", []model.SensorData{
		{Timestamp: 10, Value: 1},
		{Timestamp: 20, Value: 2},
		{Timestamp: 30, Value: 3},
		{Timestamp: 40, Value: 4},
	}, 20, 2, &model.SensorPage{Data: []model.SensorData{
		{Timestamp: 20, Value: 2},
		{Timestamp: 30, Value: 3},
	}, Cursor: ptr(int64(40))}},
	{"Cursor end", []model.SensorData{
		{Timestamp: 10, Value: 1},
		{Timestamp: 20, Value: 2},
	}, 20, 1, &model.SensorPage{Data: []model.SensorData{
		{Timestamp: 20, Value: 2},
	}, Cursor: nil}},
}

func TestSensorMemoryRepository(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewSensorMemoryRepository(tt.data)

			got, err := repo.GetPage(t.Context(), tt.cursor, tt.limit)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(got, tt.result) {
				t.Errorf("got %v, want %v", got, tt.result)
			}
		})
	}
}
