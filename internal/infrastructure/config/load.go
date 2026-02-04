package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func Load() (*Config, error) {
	if !CheckConfig() {
		if err := CreateConfigFile(); err != nil {
			return nil, err
		}
		fmt.Println("Файл настроек успешно создан! Не забудьте добавить фильтр с элементами.")
	}

	data, err := os.ReadFile("../configs/config.json")
	if err != nil {
		return nil, err
	}

	var config Config
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()

	err = decoder.Decode(&config)

	if err != nil {
		return nil, err
	}
	commandsInformation()
	changeConfig(&config)
	return &config, nil
}
