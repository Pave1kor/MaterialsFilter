package app

import (
	cfg "MaterialsFilter/internal/domain/config"
	json "MaterialsFilter/internal/infrastructure/json"
	cli "MaterialsFilter/internal/ui/cli"
	"os"
)

func Load(configPath string, inputFunc func() (string, error)) (*cfg.Config, error) {

	if !fileExist(configPath) {
		var defaultConfig cfg.Config
		if err := cli.WriteJSONUI(defaultConfig, configPath, inputFunc); err != nil {
			return nil, err
		}
	}
	obj, err := json.ReadJSON(configPath)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// Проверка: существует или конфиг файл
func fileExist(configPath string) bool {
	_, err := os.Stat(configPath)
	return err == nil
}
