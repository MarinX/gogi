package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
		return
	}

	debug, err := strconv.ParseBool(os.Getenv("APP_DEBUG"))
	if err != nil {
		log.Println("Cannot parse debug level, default to false")
	}
	if debug {
		log.SetLevel(log.DebugLevel)
	}

	obd := getOBDDevice(debug)
	if err := obd.Open(); err != nil {
		log.Error(err)
		return
	}
	defer obd.Close()

	db := getStore()
	if err := db.Open(); err != nil {
		log.Error(err)
		return
	}
	defer db.Close()

	// start the service that connects the reader and database
	service := NewService(obd, db)
	service.Run()
}

func getOBDDevice(debug bool) *OBD {
	useFake, err := strconv.ParseBool(os.Getenv("USE_FAKE"))
	if err != nil {
		log.Println("Cannot parse fake, default to false")
	}

	devicePath := os.Getenv("SERIAL_DEVICE")
	if len(devicePath) <= 0 {
		devicePath = "/dev/ttyUSB0"
	}

	return NewOBD(devicePath, debug, useFake)
}

func getStore() Store {
	layer := os.Getenv("DB_DRIVER")
	switch layer {
	case "influx":
		return NewInfluxDB(
			fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_DATABASE"),
		)
	}
	return NewInfluxDB(
		fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
	)
}
