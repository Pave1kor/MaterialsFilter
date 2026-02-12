package config

import (
	path "MaterialsFilter/internal/infrastructure/path"
	"encoding/json"
	"fmt"
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

func NewConfig() (*Config, error) {

	pathFile := path.New("..")
	configPath, err := pathFile.Config()
	if err != nil {
		return nil, err
	}

	if !fileExist(configPath) {
		createJSONFile(configPath)
	}
	obj, err := readJSON(configPath)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// Проверка: существует или конфиг файл
func fileExist(configPath string) bool {

	if _, err := os.Stat(configPath); err == nil {
		return true
	}

	return false
}

// читаем данные из json
func readJSON(configPath string) (Config, error) {
	var jsonFile Config
	fileData, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}
	err = json.Unmarshal(fileData, &jsonFile)
	if err != nil {
		return Config{}, err
	}
	return jsonFile, nil
}

// Создание нового файла конфигурация без фильтра
func createJSONFile(configPath string) error {

	fmt.Println()
	fmt.Println("Файл настроек отсутвует!")
	fmt.Println("Создание нового файла настроек.")
	pathInput := path.New("..")

	input, err := pathInput.Input()
	if err != nil {
		return err
	}
	var defaultConfig Config
	defaultConfig = Config{
		Input:   input,
		Filters: []Filter{}}

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
