package device

import (
	"fmt"
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
