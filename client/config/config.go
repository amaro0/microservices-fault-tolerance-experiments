package config

import (
	"github.com/amaro0/envloader"
	"log"
)

type ExperimentConfig struct {
	ProxyServerUrl     string `env:"FINAL_SERVER_URL" envDefault:"http://localhost:4000/proxy" validate:"url"`
	ConcurrentRequests int    `env:"CONCURRENT_REQUESTS" envDefault:"100" validate:"numeric,min=1"`
	RequestBatch       int    `env:"REQUEST_BATCH" envDefault:"100" validate:"numeric,min=1"`
	Rate               int    `env:"RATE" envDefault:"4" validate:="numeric,min=1"`
}

var experimentConfigInstance *ExperimentConfig

func GetExperimentConfig() *ExperimentConfig {
	if experimentConfigInstance == nil {
		err, conf := envloader.Load(ExperimentConfig{})

		if err != nil {
			log.Fatal("Experiment config loading failed")
		}

		experimentConfigInstance = conf.(*ExperimentConfig)
	}

	return experimentConfigInstance
}
