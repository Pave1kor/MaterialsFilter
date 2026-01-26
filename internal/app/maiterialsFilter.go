package app

import (
	cfg "MaterialsFilter/config"
	material "MaterialsFilter/internal/domain/materials"
	"log"
)

func Run(config *cfg.Config) {
	for _, file := range config.Filters {
		material := material.NewMaterials()
		if err := material.ReaderCSV(file.Input); err != nil {
			log.Fatal(err)
		}
		material.ParseMaterials()
		material.MaterialsFilter(file)
		if err := material.WriteCSV(file); err != nil {
			log.Fatal(err)
		}
	}

}
