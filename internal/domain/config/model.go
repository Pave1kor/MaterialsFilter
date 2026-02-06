package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Input   string `json:"input"`
	Filters []Filter
}

type Filter struct {
	Name   string            `json:"name"`
	Filter map[string]string `json:"filter"`
	Output string            `json:"output"`
}

func NewConfig(configPath string) (*Config, error) {

	var cfg Config
	err := UnmarshalingConfig(configPath, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func UnmarshalingConfig(configPath string, config *Config) error {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, config)
	if err != nil {
		return err
	}
	return nil
}
