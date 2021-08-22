package influxv2

import (
	"fmt"
	"net/url"

	"github.com/io-1/kuiper/internal/logger"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type InfluxV2Persistence struct {
	logger logger.Logger
	client influxdb2.Client
	org    string
	bucket string
}

func NewInfluxV2Persistence(url *url.URL, org, token string, logger logger.Logger) (*InfluxV2Persistence, error) {
	i := &InfluxV2Persistence{}

	bucket := url.Path[1:len(url.Path)]
	if bucket == "" {
		bucket = "test"
	}

	addr := fmt.Sprintf("%s://%s", url.Scheme, url.Host)
	// client := influxdb2.NewClientWithOptions(addr, token, influxdb2.DefaultOptions().SetBatchSize(20))
	client := influxdb2.NewClient(addr, token)
	i = &InfluxV2Persistence{
		logger: logger,
		client: client,
		org:    org,
		bucket: bucket,
	}

	return i, nil
}

func (i *InfluxV2Persistence) Close() {
	i.client.Close()
}
