package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Проверка: существует или конфиг файл
func CheckConfig() bool {
	_, err := os.Stat("../configs/config.json")
	return !os.IsNotExist(err)
}

// Создание нового файла конфигурация без фильтра
func CreateConfigFile() error {
	fmt.Println("Создание нового файла конфигурации.")
	var filterArr []Filter

	var defaultConfig Config
	defaultConfig = Config{
		Input:   "../data/input/inputData.csv",
		Filters: filterArr,
	}

	return MarshalingConfig(defaultConfig)
}

// Запись конфигурационного файла
func MarshalingConfig(defaultConfig Config) error {

	jsonFile, err := os.Create("../configs/config.json")
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
