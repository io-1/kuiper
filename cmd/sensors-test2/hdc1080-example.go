package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/url"
	"time"

	"github.com/io-1/kuiper/internal/logger/logruslogger"
	"github.com/io-1/kuiper/internal/sensors/persistence/influxv2"

	sensors "github.com/io-1/kuiper/internal/sensors/devicesensors"
)

const (
	org   = "io1"
	token = "aW8xOmt1aXBlcg=="
	u     = "http://root:password@172.28.1.4:8086/sensors"
)

func main() {
	mac := "1122334455FF"
	mac2 := "0022445577AA"

	ctx := context.Background()
	logger := logruslogger.NewLogrusLogger(true)

	influxUrl, err := url.Parse(u)
	if err != nil {
		logger.Fatal(err.Error())
	}

	persistence, err := influxv2.NewInfluxV2Persistence(influxUrl, org, token, logger)
	if err != nil {
		logger.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		r := rand.Intn(10)
		s := &sensors.HDC1080Measurement{
			Mac:         mac,
			Humidity:    float64(1 + r),
			Temperature: float64(2 + r),
		}
		err = persistence.CreateHDC1080Measurement(ctx, s)
		if err != nil {
			logger.Fatal(err)
		}

		s = &sensors.HDC1080Measurement{
			Mac:         mac2,
			Humidity:    float64(1 + r),
			Temperature: float64(2 + r),
		}
		err = persistence.CreateHDC1080Measurement(ctx, s)
		if err != nil {
			logger.Fatal(err)
		}
	}

	now := time.Now()
	startTime := now.Add(-time.Second * 10).UTC()
	endTime := now.Add(time.Second * 10).UTC()

	measurements, err := persistence.GetHDC1080TemperatureMeasurements(ctx, mac, startTime, endTime)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Printf("%v\n", measurements)
}
