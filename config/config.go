package config

import (
	"github.com/jinzhu/configor"
	"log"
)

type APIConfig struct {
	GRPC struct {
		Port int16 `default:"9090"`
	}
	HTTP struct {
		Port int16 `default:"9080"`
	}
}

var (
	defaultConfigFilename = "config.yml"
	configEnvPrefix       = "API"
	apiConfig             = APIConfig{}
)

func LoadConfig(configFilename string) *APIConfig {
	if configFilename == "" {
		configFilename = defaultConfigFilename
	}

	if err := configor.New(&configor.Config{
		AutoReload: false,
		ENVPrefix:  configEnvPrefix,
	}).Load(&apiConfig, configFilename); err != nil {
		log.Fatalf("failed to load configuration file `%s`: %+v", defaultConfigFilename, err)
	}
	return &apiConfig
}
