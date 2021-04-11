package config

import (
	"github.com/amaro0/envloader"
	"log"
)

type ServerConfig struct {
	Port          string `env:"PORT" envDefault:"3000" validate:"numeric"`
	GinMode       string `env:"GIN_MODE" envDefault:"debug" validate:"oneof=debug release"`
	ShouldPodFail bool   `env:"SHOULD_POD_FAIL" envDefault:"false"`
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
