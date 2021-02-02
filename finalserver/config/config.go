package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/go-playground/validator/v10"
	"log"
)

type Config struct {
	Port    string `env:"PORT" envDefault:"3002" validate:"numeric"`
	GinMode string `env:"GIN_MODE envDefault:"debug" validate:"oneof="debug release"`
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
