package influxv2

import (
	"context"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	sensors "github.com/io-1/kuiper/internal/sensors/devicesensors"
)

func (i *InfluxV2Persistence) CreateStatsMeasurement(ctx context.Context, sensor *sensors.StatsMeasurement) error {

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

	writeAPI := i.client.WriteAPIBlocking(i.org, i.bucket)
	p := influxdb2.NewPoint(
		"stats_listener",
		tags,
		fields,
		time.Now().UTC())

	err := writeAPI.WritePoint(ctx, p)
	if err != nil {
		return err
	}

	return nil
}
