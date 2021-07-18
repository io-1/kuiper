package influxpersistence

import (
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	sensors "github.com/io-1/kuiper/internal/sensors/devicesensors"
)

func (i InfluxPersistence) CreateMC38Measurement(sensor *sensors.MC38Measurement) error {
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
		"state": sensor.State,
	}

	point, err := client.NewPoint(
		"mc83_listener",
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
