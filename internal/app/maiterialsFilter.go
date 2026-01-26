package app

import (
	cfg "MaterialsFilter/internal/config"
	material "MaterialsFilter/internal/domain/materials"
	"log"
)

func Run(config *cfg.Config) {
	for _, path := range config.Paths {
		material := material.NewMaterials()
		if err := material.ReaderCSV(path.InputData); err != nil {
			log.Fatal(err)
		}
		material.ParseMaterials()
		material.MaterialsFilter(config.Filters)
		if err := material.WriteCSV(config.Filters); err != nil {
			log.Fatal(err)
		}
	}

}
