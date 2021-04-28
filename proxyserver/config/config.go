package config

import (
	"github.com/amaro0/envloader"
	"log"
)

type ProxyProtectionType string

const (
	Timeout        ProxyProtectionType = "timeout"
	CircuitBreaker ProxyProtectionType = "breaker"
)

type ServerConfig struct {
	Port                 string                `env:"PORT" envDefault:"4000" validate:"numeric"`
	GinMode              string                `env:"GIN_MODE envDefault:"debug" validate:"oneof="debug release"`
	FinalServerUrl       string                `env:"FINAL_SERVER_URL" envDefault:"http://localhost:3000/experiment" validate:"url"`
	FinalInstancesCount  int                   `env:"FINAL_INSTANCES_COUNT" envDefault:"1" validate:"number,min=1"`
	MetricsServerUrl     string                `env:"METRICS_SERVER_URL" envDefault:"http://localhost:2000" validate:"url"`
	ProxyProtectionTypes []ProxyProtectionType `env:"PROXY_PROTECTION_TYPE" envDefault:""`
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

func (c *ServerConfig) ProtectionIncluded(p ProxyProtectionType) bool {
	for _, item := range serverConfigInstance.ProxyProtectionTypes {
		if item == p {
			return true
		}
	}
	return false
}
