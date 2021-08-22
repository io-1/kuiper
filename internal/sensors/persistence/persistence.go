package persistence

import (
	"context"
	"time"

	sensors "github.com/io-1/kuiper/internal/sensors/devicesensors"
)

type BMP280Sensor interface {
	CreateBMP280Measurement(ctx context.Context, sensor *sensors.BMP280Measurement) error
}

type DHT22Sensor interface {
	CreateDHT22Measurement(ctx context.Context, sensor *sensors.DHT22Measurement) error
}

type HDC1080Sensor interface {
	CreateHDC1080Measurement(ctx context.Context, sensor *sensors.HDC1080Measurement) error
	GetHDC1080TemperatureMeasurements(ctx context.Context, mac string, startTime, endTime time.Time) (sensors.HDC1080TemperatureMeasurements, error)
	GetHDC1080HumidityMeasurements(ctx context.Context, mac string, startTime, endTime time.Time) (sensors.HDC1080HumidityMeasurements, error)
}

type MC38Sensor interface {
	CreateMC38Measurement(ctx context.Context, sensor *sensors.MC38Measurement) error
}

type HCSR501Sensor interface {
	CreateHCSR501Measurement(ctx context.Context, sensor *sensors.HCSR501Measurement) error
}

type BH1750Sensor interface {
	CreateBH1750Measurement(ctx context.Context, sensor *sensors.BH1750Measurement) error
}

type Stats interface {
	CreateStatsMeasurement(ctx context.Context, sensor *sensors.StatsMeasurement) error
}

type Voltage interface {
	CreateVoltageMeasurement(ctx context.Context, sensor *sensors.VoltageMeasurement) error
}

type Keypad interface {
	CreateKeypadMeasurement(ctx context.Context, sensor *sensors.KeypadMeasurement) error
}

type Persistence interface {
	BMP280Sensor
	DHT22Sensor
	HDC1080Sensor
	MC38Sensor
	HCSR501Sensor
	BH1750Sensor
	Stats
	Voltage
	Keypad
}
