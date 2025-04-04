package main

import (
	"context"
	"log"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

func main() {
	token := "ytMTmUT7bivfL8HxR8IlXP8NZTZedLTConSlh0rrmfF-pQvfIhpnsN3kK62wQ1afQECkmG0nlHRh2Kd5DbRZYw=="
	url := "http://localhost:8086"
	client := influxdb2.NewClient(url, token)


	org := "MPT"
	bucket := "metrics"
	writeAPI := client.WriteAPIBlocking(org, bucket)
	for value := 0; value < 5; value++ {
		tags := map[string]string{
			"tagname1": "tagvalue1",
		}
		fields := map[string]interface{}{
			"field1": value,
		}
		point := write.NewPoint("measurement1", tags, fields, time.Now())
		time.Sleep(1 * time.Second) // separate points by 1 second

		if err := writeAPI.WritePoint(context.Background(), point); err != nil {
			log.Fatal(err)
		}
	}
}