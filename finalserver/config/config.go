package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/go-playground/validator/v10"
	"log"
)

type Config struct {
	Port    string `env:"PORT" envDefault:"3002" validate:"hostname_port"`
	GinMode string `'env:"GIN_MODE' envDefault:"debug" validate:"oneof="debug release"`
}

func NewConfig() *Config {
	c := Config{}

	if err := env.Parse(&c); err != nil {
		log.Fatal("Env parse error! ", err)
	}

	validate := validator.New()
	if err := validate.Struct(c); err != nil {
		log.Fatal("Env validation error! ", err.Error())
	}

	return &c
}
