package json

import (
	cfg "MaterialsFilter/internal/domain/config"
	"encoding/json"
	"log"
	"os"
)

// Сохранение данных в конфигурационный файл
func WriteJSON(config *cfg.Config) {

	jsonData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(config.ConfigPathFile, jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
