package device

import (
	"github.com/smarthut/agent/model/device/laurent112"
	"github.com/smarthut/agent/model/device/megad328"
)

// Factory ...
type Factory func(host, password string) (Device, error)

func newMegaD328Device(host, password string) (Device, error) {
	return megad328.New(host, password), nil
}

func newLaurent112Device(host, password string) (Device, error) {
	return laurent112.New(host), nil
}

// Drivers ...
var Drivers = map[string]Factory{
	"megad328":   newMegaD328Device,
	"laurent112": newLaurent112Device,
}
