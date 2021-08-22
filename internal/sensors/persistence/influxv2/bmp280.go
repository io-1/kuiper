package influxv2

import (
	"context"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	sensors "github.com/io-1/kuiper/internal/sensors/devicesensors"
)

func (i *InfluxV2Persistence) CreateBMP280Measurement(ctx context.Context, sensor *sensors.BMP280Measurement) error {
	pressureFloat, err := sensor.GetPressureFloat()
	if err != nil {
		return err
	}

	temperatureFloat, err := sensor.GetTemperatureFloat()
	if err != nil {
		return err
	}

	// indexed
	tags := map[string]string{
		"mac": sensor.Mac,
	}

	// not indexed
	fields := map[string]interface{}{
		"pressure": pressureFloat,
		"temp":     temperatureFloat,
	}

	writeAPI := i.client.WriteAPIBlocking(i.org, i.bucket)

	p := influxdb2.NewPoint(
		"bmp280_listener",
		tags,
		fields,
		time.Now().UTC())

	err = writeAPI.WritePoint(ctx, p)
	if err != nil {
		return err
	}

	return nil
}
