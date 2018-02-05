package device

import (
	"fmt"
	"log"
	"time"

	"github.com/caarlos0/env"
)

var (
	currentDevice Device
	config        Config
)

// Device provides interface for accessing the device
type Device interface {
	UpdateSockets() error            // updates all sockets
	Get(id int) (interface{}, error) // get information from socket
	Set(id, value int) error         // sends information to socket
	Len() int                        // get sockets num
}

// New creates the specified device
func New(driver, host, password string) (Device, error) {
	factory, ok := Drivers[driver]
	if !ok {
		return nil, fmt.Errorf("agent: there are no driver for %s device", driver)
	}

	dev, err := factory(host, password)
	if err != nil {
		return nil, err
	}

	return dev, nil
}

// Create creates the device
func Create() error {
	err := env.Parse(&config)
	if err != nil {
		return err
	}

	currentDevice, err = New(config.Driver, config.Host, config.Password)

	if err != nil {
		return err
	}

	go startPolling()

	return nil
}

// Get returns the device
func Get() Device {
	return currentDevice
}

func startPolling() {
	for {
		if err := currentDevice.UpdateSockets(); err != nil {
			log.Println(err)
		}
		<-time.After(config.PollingTime * time.Second)
	}
}
