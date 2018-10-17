package main

import (
	"time"

	"github.com/rzetterberg/elmobd"
	log "github.com/sirupsen/logrus"
)

// OBD holds information about the device.
// Also holds channels for data query
type OBD struct {
	dev    *elmobd.Device
	device string
	debug  bool
	fake   bool
	closed bool
	chData chan Data
}

// NewOBD creates a new device
func NewOBD(device string, debug bool, fake bool) *OBD {
	return &OBD{
		device: device,
		debug:  debug,
		chData: make(chan Data, 10),
	}
}

// Open opens serial or bluetooth device
func (o *OBD) Open() error {
	var err error
	if o.fake {
		o.dev, err = elmobd.NewTestDevice(o.device, o.debug)
	} else {
		o.dev, err = elmobd.NewDevice(o.device, o.debug)
	}
	return err
}

// Run device and starts to read from OBD
func (o *OBD) Run() error {
	supported, err := o.dev.CheckSupportedCommands()
	if err != nil {
		return err
	}
	commands := supported.FilterSupported(elmobd.GetSensorCommands())
	for {
		if o.closed {
			close(o.chData)
			break
		}

		results, err := o.dev.RunManyOBDCommands(commands)
		if err != nil {
			log.Error(err)
		}

		for i := range results {
			o.chData <- Data{
				Key:   results[i].Key(),
				Value: results[i].ValueAsLit(),
			}
		}

		time.Sleep(time.Second)
	}

	return nil
}

// Close is ping to exit loop and close channels
func (o *OBD) Close() {
	o.closed = true
}
