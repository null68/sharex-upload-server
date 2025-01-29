package config

import (
	"errors"
	"os"

	"github.com/goccy/go-json"
)

type Config struct {
	Port   int    `json:"port"`
	ApiKey string `json:"api_key"`
}

func DefaultConfig() *Config {
	return &Config{
		Port:   8080,
		ApiKey: "default_api_key",
	}
}

func CreateConfigFile(filePath string) error {
	var defaultConfig *Config = DefaultConfig()

	file, err := os.Create(filePath)
	if err != nil {
		return errors.New(err.Error())
	}
	defer file.Close()

	encoder := json.NewEncoder(file);
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(defaultConfig); err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func LoadConfig(filePath string) (*Config, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if err := CreateConfigFile(filePath); err != nil {
			return nil, errors.New(err.Error())
		}
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, errors.New(err.Error())
	} 

	return &cfg, nil
}