package json

import (
	cfg "MaterialsFilter/internal/domain/config"
	"encoding/json"
	"os"
)

// читаем данные из json
func ReadJSON(configPath string) (cfg.Config, error) {
	var jsonFile cfg.Config
	fileData, err := os.ReadFile(configPath)
	if err != nil {
		return cfg.Config{}, err
	}
	err = json.Unmarshal(fileData, &jsonFile)
	if err != nil {
		return cfg.Config{}, err
	}
	return jsonFile, nil
}
