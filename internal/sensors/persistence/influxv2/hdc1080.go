package influxv2

import (
	"context"
	"fmt"
	"time"

	"github.com/araddon/dateparse"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	sensors "github.com/io-1/kuiper/internal/sensors/devicesensors"
)

const (
	format = "2006-01-02T15:04:05Z"
)

func (i *InfluxV2Persistence) CreateHDC1080Measurement(ctx context.Context, sensor *sensors.HDC1080Measurement) error {

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
		"hdc1080_listener",
		tags,
		fields,
		time.Now().UTC())

	err := writeAPI.WritePoint(ctx, p)
	if err != nil {
		return err
	}

	return nil
}

func (i *InfluxV2Persistence) GetHDC1080TemperatureMeasurements(ctx context.Context, mac string, startTime, endTime time.Time) (sensors.HDC1080TemperatureMeasurements, error) {
	var (
		start                          = startTime.Format(format)
		stop                           = endTime.Format(format)
		hdc1080TemperatureMeasurements sensors.HDC1080TemperatureMeasurements
	)

	query := fmt.Sprintf(
		`from(bucket:"%s")
			|> range(start: %s, stop: %s) 
			|> filter(fn: (r) => 
				r._measurement == "hdc1080_listener" 
					and r.mac == "%s" 
					and r._field == "temp")`, i.bucket, start, stop, mac)

	queryAPI := i.client.QueryAPI(i.org)
	result, err := queryAPI.Query(ctx, query)
	if err == nil {
		for result.Next() {
			values := result.Record().Values()
			t := fmt.Sprintf("%s", values["_time"])
			temp := values["_value"].(float64)

			timestamp, err := dateparse.ParseAny(t)
			if err != nil {
				i.logger.Error(err)
			} else {
				m := sensors.HDC1080TemperatureMeasurement{
					Timestamp:   timestamp,
					Temperature: temp,
				}

				hdc1080TemperatureMeasurements.TemperatureMeasurements = append(hdc1080TemperatureMeasurements.TemperatureMeasurements, m)
			}
		}
		if result.Err() != nil {
			return hdc1080TemperatureMeasurements, result.Err()
		}
	}

	hdc1080TemperatureMeasurements.Mac = mac

	return hdc1080TemperatureMeasurements, nil
}

func (i *InfluxV2Persistence) GetHDC1080HumidityMeasurements(ctx context.Context, mac string, startTime, endTime time.Time) (sensors.HDC1080HumidityMeasurements, error) {
	var (
		start                       = startTime.Format(format)
		stop                        = endTime.Format(format)
		hdc1080HumidityMeasurements sensors.HDC1080HumidityMeasurements
	)

	query := fmt.Sprintf(
		`from(bucket:"%s")
			|> range(start: %s, stop: %s) 
			|> filter(fn: (r) => 
				r._measurement == "hdc1080_listener" 
					and r.mac == "%s" 
					and r._field == "humidity")`, i.bucket, start, stop, mac)

	queryAPI := i.client.QueryAPI(i.org)
	result, err := queryAPI.Query(ctx, query)
	if err == nil {
		for result.Next() {
			values := result.Record().Values()
			t := fmt.Sprintf("%s", values["_time"])
			humidity := values["_value"].(float64)

			timestamp, err := dateparse.ParseAny(t)
			if err != nil {
				i.logger.Error(err)
			} else {
				m := sensors.HDC1080HumidityMeasurement{
					Timestamp: timestamp,
					Humidity:  humidity,
				}

				hdc1080HumidityMeasurements.HumidityMeasurements = append(hdc1080HumidityMeasurements.HumidityMeasurements, m)
			}
		}
		if result.Err() != nil {
			return hdc1080HumidityMeasurements, result.Err()
		}
	}

	hdc1080HumidityMeasurements.Mac = mac

	return hdc1080HumidityMeasurements, nil
}
