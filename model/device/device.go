package device

import (
	"fmt"
	"log"
	"time"
)

var (
	currentDevice Device
	pollingTime   time.Duration
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
func Create(c Configuration) error {
	var err error
	currentDevice, err = New(c.Device.Driver, c.Device.Host, c.Device.Password)
	if err != nil {
		return err
	}

	pollingTime = c.Device.PollingTime

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
		<-time.After(pollingTime)
	}
}
