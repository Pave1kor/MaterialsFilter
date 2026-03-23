package json

import (
	cfg "MaterialsFilter/internal/domain/config"
	"encoding/json"
	"log"
	"os"
)

// Получение настроек из файла конфигурации
func ReadJSON(config *cfg.Config) {
	fileData, err := os.ReadFile(config.ConfigPathFile)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(fileData, &config)
	if err != nil {
		log.Fatal(err)
	}
}
