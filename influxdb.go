package main

import (
	"time"

	"github.com/influxdata/influxdb/client/v2"
	log "github.com/sirupsen/logrus"
)

// InfluxDB is our storage layer
type InfluxDB struct {
	username    string
	password    string
	db          string
	addr        string
	client      client.Client
	batchPoints *client.BatchPoints
}

// NewInfluxDB creates new object
func NewInfluxDB(addr, username, password, db string) *InfluxDB {
	return &InfluxDB{
		db:       db,
		addr:     addr,
		username: username,
		password: password,
	}
}

// Open connection to influxdb
func (i *InfluxDB) Open() error {
	log.Debug("opening db connection")
	var err error
	i.client, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     i.addr,
		Username: i.username,
		Password: i.password,
	})
	return err
}

// Insert into infludb database
func (i *InfluxDB) Insert(fields map[string]interface{}) error {
	log.Debugf("Fields: %+v\n", fields)
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database: i.db,
	})
	if err != nil {
		return err
	}

	point, err := client.NewPoint("obd", nil, fields, time.Now())
	if err != nil {
		return err
	}
	bp.AddPoint(point)

	return i.client.Write(bp)
}

// Close the connection to influxdb
func (i *InfluxDB) Close() error {
	log.Debug("closing db connection")
	return i.client.Close()
}
