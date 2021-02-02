package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/go-playground/validator/v10"
	"log"
)

type Config struct {
	FinalServerUrl     string `env:"FINAL_SERVER_URL" envDefault:"http://localhost:3002/experiment" validate:"url"`
	ConcurrentRequests int    `env:"CONCURRENT_REQUESTS" envDefault:"100" validate:"numeric"`
	RequestBatch       int    `env:"CONCURRENT_REQUESTS" envDefault:"100" validate:"numeric"`
}

var globalConf *Config

func newConfig() *Config {
	globalConf = &Config{}

	if err := env.Parse(globalConf); err != nil {
		log.Fatal("Env parse error! ", err)
	}

	validate := validator.New()
	if err := validate.Struct(globalConf); err != nil {
		log.Fatal("Env validation error! ", err.Error())
	}

	return globalConf
}

func GetConfig() *Config {
	if globalConf == nil {
		return newConfig()
	}

	return globalConf
}
