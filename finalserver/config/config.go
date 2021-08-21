package config

import (
	"github.com/amaro0/envloader"
	"log"
)

type ErrorType string
type Mode string

const (
	TimeoutError   ErrorType = "timeout"
	UnhandledError ErrorType = "unhandled"
	RandomizedMode Mode      = "randomizedMode"
	FailAfterMode  Mode      = "failAfterMode"
	SpikeMode      Mode      = "spikeMode"
	NoErrorMode    Mode      = "noErrorMode"
)

type ServerConfig struct {
	Port                  string    `env:"PORT" envDefault:"3000" validate:"numeric"`
	GinMode               string    `env:"GIN_MODE" envDefault:"debug" validate:"oneof=debug release"`
	ShouldServerFail      bool      `env:"SHOULD_SERVER_FAIL" envDefault:"false"`
	ShouldServerFailAfter int       `env:"SHOULD_SERVER_FAIL_AFTER" envDefault:"5"`
	HashSalt              int       `env:"HASH_SALT" envDefault:"10"`
	ErrorRatio            int       `env:"ERROR_RATIO" envDefault:"25" validate:"min=0,max=100"`
	ErrorType             ErrorType `env:"ERROR_TYPE" envDefault:"unhandled" validate:"oneof=timeout unhandled"`
	TimeoutLengthInS      int       `env:"TIMEOUT_LENGTH" envDefault:"30" validate:"min=0"`
	Mode                  Mode      `env:"MODE" envDefault:"noErrorMode" validate:"oneof=noErrorMode randomizedMode failAfterMode spikeMode"`
	FailAfterTimeInS      int       `env:"FAIL_AFTER_TIME" envDefault:"2" validate:"min=0"`
	FailDurationTimeInS   int       `env:"FAIL_DURATION_TIME" envDefault:"2" validate:"min=0"`
}

var serverConfigInstance *ServerConfig

func GetServerConfig() *ServerConfig {
	if serverConfigInstance == nil {
		err, conf := envloader.Load(ServerConfig{})

		if err != nil {
			log.Fatal("Server config loading failed")
		}

		serverConfigInstance = conf.(*ServerConfig)
	}

	return serverConfigInstance
}

func (c *ServerConfig) IsRandomizedMode() bool {
	return c.Mode == RandomizedMode
}

func (c *ServerConfig) IsFailAfterMode() bool {
	return c.Mode == RandomizedMode
}

func (c *ServerConfig) IsSpikeMode() bool {
	return c.Mode == SpikeMode
}
