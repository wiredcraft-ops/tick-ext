package tick

import (
	"log"
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

const (
	// TODO: 16/11/15 flag
	influxdbAddress  = "http://127.0.0.1:8086"
	influxdbDatabase = "wcl"
	influxdbUsername = "admin"
	influxdbPassword = "admin"
)

var (
	ct  client.Client
	err error
)

func init() {

	ct, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     influxdbAddress,
		Username: influxdbUsername,
		Password: influxdbPassword,
	})
	if err != nil {
		log.Fatalf("E! %v\n", err)
	}
}

func Store(uuid string) error {

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  influxdbDatabase,
		Precision: "s",
	})
	if err != nil {
		return err
	}

	tags := map[string]string{"uuid": uuid}
	fields := map[string]interface{}{}
	pt, err := client.NewPoint("wcl", tags, fields, time.Now())
	if err != nil {
		return err
	}
	bp.AddPoint(pt)

	return ct.Write(bp)
}
