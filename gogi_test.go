package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestInfluxDB(t *testing.T) {

	db := NewInfluxDB("http://localhost:8086", "car")
	if err := db.Open(); err != nil {
		t.Error(err)
		return
	}
	defer db.Close()

	j := 0
	for {
		if j > 10 {
			break
		}

		if err := db.Insert("obd", map[string]interface{}{"EngineRPM": rand.Intn(100)}); err != nil {
			t.Error(err)
			break
		}

		j++
		time.Sleep(time.Second)
	}

}
