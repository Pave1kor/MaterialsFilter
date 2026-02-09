package config

import (
	path "MaterialsFilter/internal/infrastructure/path"
	"encoding/json"
	"fmt"
	"os"
)

// Проверка: существует или конфиг файл
func CheckConfig(configPath string) error {
	_, err := os.Stat(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Конфигурационный файл не существует!")
			if err := createConfigFile(configPath); err != nil {
				return err
			}

		}
	}

	return nil
}

// Создание нового файла конфигурация без фильтра
func createConfigFile(configPath string) error {
	fmt.Println("Создание нового файла конфигурации.")
	pathInput := path.New("..")

	input, err := pathInput.Input()
	if err != nil {
		return err
	}

	var defaultConfig Config
	defaultConfig = Config{
		Input:   input,
		Filters: []Filter{},
	}

	jsonFile, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	jsonData, err := json.MarshalIndent(defaultConfig, "", "  ")
	if err != nil {
		return err
	}
	_, err = jsonFile.Write(jsonData)
	return err
}
