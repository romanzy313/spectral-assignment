package repository

import (
	"bufio"
	"context"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/romanzy313/spectral-assignment/pkg/server/model"
)

func ReadCsvData(fileName string) ([]model.SensorData, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var result []model.SensorData

	scanner := bufio.NewScanner(file)

	// skip first line
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()

		// split by ","
		parts := strings.Split(line, ",")

		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid format: %s", line)
		}

		time, err := time.Parse(time.DateTime, parts[0])
		if err != nil {
			return nil, fmt.Errorf("failed to parse time: %w", err)
		}
		timestamp := time.Unix()

		value, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			// TODO: this should not throw, but go on and log it maybe?
			return nil, fmt.Errorf("failed to parse value: %w", err)
		}

		// sneaky :)
		if math.IsNaN(value) || math.IsInf(value, 0) {
			continue
		}

		result = append(result, model.SensorData{
			Timestamp: timestamp,
			Value:     value,
		})
	}
	err = scanner.Err()

	if err != nil {
		return nil, scanner.Err()
	}

	return result, nil
}

type SensorMemoryRepository struct {
	SensorRepository

	data []model.SensorData
}

func NewSensorMemoryRepository(data []model.SensorData) *SensorMemoryRepository {
	return &SensorMemoryRepository{
		data: data,
	}
}

func (d *SensorMemoryRepository) GetPage(ctx context.Context, cursor int64, limit int32) (*model.SensorPage, error) {
	// find the index of where cursor is pointing at
	start := 0
	for i := 0; i < len(d.data); i++ {
		if d.data[i].Timestamp >= cursor {
			start = i
			break
		}
	}

	end := min(start+int(limit), len(d.data))

	var nextCursor *int64
	if end < len(d.data) {
		nextCursor = &d.data[end].Timestamp
	}

	// returning all data for now
	return &model.SensorPage{
		Data:   d.data[start:end],
		Cursor: nextCursor,
	}, nil
}
