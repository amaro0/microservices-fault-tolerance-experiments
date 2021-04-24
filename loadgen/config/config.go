package config

import (
	"github.com/amaro0/envloader"
	"log"
)

type ExperimentConfig struct {
	ProxyServerUrl      string `env:"PROXY_SERVER_URL" envDefault:"http://localhost:4000/proxy" validate:"url"`
	ProxyInstancesCount int    `env:"PROXY_INSTANCES_COUNT" envDefault:"1" validate:"numeric,min=1"`
	RequestBatch        int    `env:"REQUEST_BATCH" envDefault:"100" validate:"numeric,min=1"`
	Rate                int    `env:"RATE" envDefault:"4" validate:="numeric,min=1"`
	MetricsServerUrl    string `env:"METRICS_SERVER_URL" envDefault:"http://localhost:2000" validate:"url"`
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
