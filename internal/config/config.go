package config

import (
	"os"

	"gopkg.in/gcfg.v1"
)

func Init() (*Config, error) {
	cfg = &Config{}

	configFilePath := "../../files/config.ini"

	config, err := os.ReadFile(configFilePath)
	if err != nil {
		return cfg, err
	}

	err = gcfg.ReadStringInto(cfg, string(config))
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

// GetConfig returns config object
func GetConfig() *Config {
	return cfg
}
