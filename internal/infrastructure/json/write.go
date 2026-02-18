package json

import (
	cfg "MaterialsFilter/internal/domain/config"
	"encoding/json"
	"os"
)

// Создание нового файла конфигурация без фильтра
func WriteJSON(configPath string, config cfg.Config) error {

	jsonData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(configPath, jsonData, 0644)
	return err
}
