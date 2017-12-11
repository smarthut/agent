package main

// DeviceConfig holds device level config
type DeviceConfig struct {
	Driver   string `env:"DEVICE_DRIVER,required"`
	Host     string `env:"DEVICE_HOST,required"`
	Password string `env:"DEVICE_PASS"`
}
