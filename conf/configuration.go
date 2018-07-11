package conf

import "time"

// Configuration hols the configuration.
type Configuration struct {
	Device struct {
		Driver      string `required:"true"`
		Host        string `required:"true"`
		Password    string
		PollingTime time.Duration `split_words:"true" default:"5s"`
	}
	Host string
	Port int `envconfig:"PORT" default:"8080"`
}
