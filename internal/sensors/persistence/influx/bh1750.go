package influx

import (
	"context"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	sensors "github.com/io-1/kuiper/internal/sensors/devicesensors"
)

func (i InfluxPersistence) CreateBH1750Measurement(ctx context.Context, sensor *sensors.BH1750Measurement) error {
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
		"intensity": sensor.Intensity,
	}

	point, err := client.NewPoint(
		"bh1750_listener",
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
