package influx

import (
	"context"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	sensors "github.com/io-1/kuiper/internal/sensors/devicesensors"
)

func (i InfluxPersistence) CreateStatsMeasurement(ctx context.Context, sensor *sensors.StatsMeasurement) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  i.database,
		Precision: "s",
	})
	if err != nil {
		return err
	}

	// indexed
	tags := map[string]string{
		"mac": sensor.Mac,
	}

	// not indexed
	fields := map[string]interface{}{
		"voltage": sensor.Voltage,
		"connect": sensor.ConnectionTime,
		"rssi":    sensor.Rssi,
	}

	point, err := client.NewPoint(
		"stats_listener",
		tags,
		fields,
		time.Now().UTC(),
	)

	bp.AddPoint(point)

	err = i.client.Write(bp)
	if err != nil {
		return err
	}

	return nil
}
