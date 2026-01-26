package config

import (
	"bytes"
	"encoding/json"
	"os"
)

func Load() (*Config, error) {
	data, err := os.ReadFile("../internal/config/config.json")
	if err != nil {
		return nil, err
	}
	var config Config

	// err = json.Unmarshal(data, &config)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()

	err = decoder.Decode(&config)

	if err != nil {
		return nil, err
	}
	return &config, nil
}
