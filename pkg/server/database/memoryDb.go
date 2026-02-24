package database

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReadCsvData(fileName string) ([]*SensorReading, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var result []*SensorReading

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
		value, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse value: %w", err)
		}

		result = append(result, &SensorReading{
			Timestamp:  time,
			MeterUsage: value,
		})
	}
	err = scanner.Err()

	if err != nil {
		return nil, scanner.Err()
	}

	return result, nil
}

type MemoryDatabase struct {
	data []*SensorReading
}

func NewMemoryDatabase(data []*SensorReading) *MemoryDatabase {
	return &MemoryDatabase{
		data: data,
	}
}

func (d *MemoryDatabase) GetSensorReadings(ctx context.Context, req ReadingRequest) ([]*SensorReading, error) {
	// returning all data for now
	return d.data, nil
}
