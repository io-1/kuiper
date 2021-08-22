package influx

import (
	"context"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	sensors "github.com/io-1/kuiper/internal/sensors/devicesensors"
)

func (i InfluxPersistence) CreateBMP280Measurement(ctx context.Context, sensor *sensors.BMP280Measurement) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  i.database,
		Precision: "s",
	})
	if err != nil {
		return err
	}

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

	point, err := client.NewPoint(
		"bmp280_listener",
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
