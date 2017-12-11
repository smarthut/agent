package device

import (
	"fmt"

	"github.com/smarthut/agent/device/laurent112"
	"github.com/smarthut/agent/device/megad328"
)

// Device provides interface for accessing the device
type Device interface {
	UpdateSockets() error            // updates all sockets
	Get(id int) (interface{}, error) // get information from socket
	Set(id, value int) error         // sends information to socket
	Len() int                        // get sockets num
}

// Factory ...
type Factory func(host, password string) (Device, error)

func newMegaD328Device(host, password string) (Device, error) {
	return megad328.New(host, password), nil
}

func newLaurent112Device(host, password string) (Device, error) {
	return laurent112.New(host), nil
}

// NewDevice ...
func NewDevice(driver, host, password string) (Device, error) {
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

// Drivers ...
var Drivers = map[string]Factory{
	"megad328":   newMegaD328Device,
	"laurent112": newLaurent112Device,
}
