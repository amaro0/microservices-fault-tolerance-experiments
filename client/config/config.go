package config

import (
	"github.com/amaro0/envloader"
	"log"
)

type ExperimentConfig struct {
	FinalServerUrl     string `env:"FINAL_SERVER_URL" envDefault:"http://localhost:3002/experiment" validate:"url"`
	ConcurrentRequests int    `env:"CONCURRENT_REQUESTS" envDefault:"100" validate:"numeric"`
	RequestBatch       int    `env:"REQUEST_BATCH" envDefault:"100" validate:"numeric"`
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
