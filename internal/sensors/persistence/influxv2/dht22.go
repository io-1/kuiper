package influxv2

import (
	"context"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	sensors "github.com/io-1/kuiper/internal/sensors/devicesensors"
)

func (i *InfluxV2Persistence) CreateDHT22Measurement(ctx context.Context, sensor *sensors.DHT22Measurement) error {

	// indexed
	tags := map[string]string{
		"mac": sensor.Mac,
	}

	// not indexed
	fields := map[string]interface{}{
		"humidity": sensor.Humidity,
		"temp":     sensor.Temperature,
	}

	writeAPI := i.client.WriteAPIBlocking(i.org, i.bucket)
	p := influxdb2.NewPoint(
		"dht22_listener",
		tags,
		fields,
		time.Now().UTC())

	err := writeAPI.WritePoint(ctx, p)
	if err != nil {
		return err
	}

	return nil
}
