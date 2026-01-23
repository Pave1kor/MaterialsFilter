package app

import (
	cfg "MaterialsFilter/internal/config"
	material "MaterialsFilter/internal/domain/materials"
	"log"
)

func Run(config *cfg.Config) {

	material := material.NewMaterials()
	if err := material.ReaderCSV(config.Paths.InputData); err != nil {
		log.Fatal(err)
	}
	material.ParseMaterials()
	material.MaterialsFilter(*config)
	if err := material.WriteCSV(*config); err != nil {
		log.Fatal(err)
	}
}
