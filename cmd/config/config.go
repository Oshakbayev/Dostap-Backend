package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	HTTPPort string `json:"HTTPPort"`
	DBAddr   string `json:"DBAddr"`
	DBDriver string `json:"DBDriver"`
}

func CreateConfig() *Config {
	return &Config{}
}

func ReadConfig(configFilePath string, config *Config) error {
	
	configJSONData, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(configJSONData, config); err != nil {
		return err
	}
	
	if port := os.Getenv("PORT"); port != "" {
		config.HTTPPort = port
	}

	return nil
}
