package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Проверка: существует или конфиг файл
func CheckConfig(configPath string) error {
	_, err := os.Stat(configPath)
	if err != nil {
		return err
	}

	if os.IsNotExist(err) {
		fmt.Println("Конфигурационный файл не существует. Создаю новый файл...")
		if err := createConfigFile(configPath); err != nil {
			return err
		}

	}
	return nil
}

// Создание нового файла конфигурация без фильтра
func createConfigFile(configPath string) error {
	fmt.Println("Создание нового файла конфигурации.")

	var defaultConfig Config
	defaultConfig = Config{
		Input:   "",
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
