package config

import (
	"github.com/amaro0/envloader"
	"log"
)

type serverConfig struct {
	Port           string `env:"PORT" envDefault:"4000" validate:"numeric"`
	GinMode        string `env:"GIN_MODE envDefault:"debug" validate:"oneof="debug release"`
	FinalServerUrl string `env:"FINAL_SERVER_URL" envDefault:"http://localhost:3000/experiment" validate:"url"`
}

var serverConfigInstance *serverConfig

func GetServerConfig() *serverConfig {
	if serverConfigInstance == nil {
		err, conf := envloader.Load(serverConfig{})

		if err != nil {
			log.Fatal("Server config loading failed")
		}

		serverConfigInstance = conf.(*serverConfig)
	}

	return serverConfigInstance
}
