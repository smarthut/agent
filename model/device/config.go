package device

import "time"

// Config holds device level config
type Config struct {
	Driver      string        `env:"DEVICE_DRIVER,required"`
	Host        string        `env:"DEVICE_HOST,required"`
	Password    string        `env:"DEVICE_PASS"`
	PollingTime time.Duration `env:"DEVICE_POLLING_TIME",envDefault:"5"`
}
