package config

import (
	"encoding/json"
	"os"
)

func SaveConfig(config *Config, configPath string) error {
	err := os.Remove(configPath)
	if err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(configPath, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}
