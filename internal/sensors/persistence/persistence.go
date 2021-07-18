package persistence

import (
	"time"

	sensors "github.com/io-1/kuiper/internal/sensors/devicesensors"
)

type BMP280Sensor interface {
	CreateBMP280Measurement(sensor *sensors.BMP280Measurement) error
}

type DHT22Sensor interface {
	CreateDHT22Measurement(sensor *sensors.DHT22Measurement) error
}

type HDC1080Sensor interface {
	CreateHDC1080Measurement(sensor *sensors.HDC1080Measurement) error
	GetHDC1080TemperatureMeasurements(mac string, startTime, endTime time.Time) (sensors.HDC1080TemperatureMeasurements, error)
	GetHDC1080HumidityMeasurements(mac string, startTime, endTime time.Time) (sensors.HDC1080HumidityMeasurements, error)
}

type MC38Sensor interface {
	CreateMC38Measurement(sensor *sensors.MC38Measurement) error
}

type HCSR501Sensor interface {
	CreateHCSR501Measurement(sensor *sensors.HCSR501Measurement) error
}

type BH1750Sensor interface {
	CreateBH1750Measurement(sensor *sensors.BH1750Measurement) error
}

type Stats interface {
	CreateStatsMeasurement(sensor *sensors.StatsMeasurement) error
}

type Voltage interface {
	CreateVoltageMeasurement(sensor *sensors.VoltageMeasurement) error
}

type Keypad interface {
	CreateKeypadMeasurement(sensor *sensors.KeypadMeasurement) error
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
