package main

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

// Service is responsible to controll the data flow from OBD to storage layer
type Service struct {
	store Store
	obd   *OBD
	kill  chan os.Signal
}

// NewService creates a new service with OBD and store interface
func NewService(obd *OBD, store Store) *Service {
	return &Service{
		obd:   obd,
		store: store,
		kill:  make(chan os.Signal),
	}
}

// Run reads from channel and push it to the store
func (s *Service) Run() {
	signal.Notify(s.kill, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case v := <-s.obd.chData:
			log.Debugf("Received %s with value %s", v.Key, v.Value)
			s.chanInsert(map[string]interface{}{v.Key: v.Value})
			break
		case <-s.kill:
			log.Debug("Kill signal received")
			return
		}
	}
}

func (s *Service) chanInsert(fields map[string]interface{}) {
	if err := s.store.Insert(fields); err != nil {
		log.Error(err)
		// we have an error inserting into db - shutdown
		s.kill <- os.Interrupt
	}
}
