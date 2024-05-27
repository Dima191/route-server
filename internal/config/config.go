package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Config struct {
	ConnectionStr string `env:"CONNECTION_STRING" env-required:"true"`
	AppPort       int    `env:"APP_PORT" env-required:"true"`
}

var once sync.Once

func New(configPath string) (*Config, error) {
	var err error
	cfg := &Config{}

	once.Do(func() {
		err = cleanenv.ReadConfig(configPath, cfg)
	})

	return cfg, err
}
