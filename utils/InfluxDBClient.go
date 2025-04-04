package utils

import (
	"context"
	"time"

	"STDE_proj/configs"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

var InfluxClient influxdb2.Client

func InitInfluxDBClient() error {
	InfluxClient = influxdb2.NewClient(configs.AppConfig.Influxdb.URL, configs.AppConfig.Influxdb.Token)
	return nil
}

func WriteMetric(measurement string, tags map[string]string, fields map[string]interface{}) error {
	org := configs.AppConfig.Influxdb.Org
	bucket := configs.AppConfig.Influxdb.Bucket
	writeAPI := InfluxClient.WriteAPIBlocking(org, bucket)

	point := write.NewPoint(measurement, tags, fields, time.Now())
	return writeAPI.WritePoint(context.Background(), point)
}
