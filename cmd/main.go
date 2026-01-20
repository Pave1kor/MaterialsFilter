package main

import (
	materials "MaterialsFilter/internal/materials"
	"log"
)

func main() {
	material := materials.NewMaterials()
	if err := material.ReaderCSV(); err != nil {
		log.Fatal(err)
	}
	material.ParseMaterials()
	material.MaterialsFilterHeusler()
	material.MaterialsFilterChalcogenide()
	if err := material.WriteCSVHeusler(); err != nil {
		log.Fatal(err)
	}
	if err := material.WriteCSVChalcogenide(); err != nil {
		log.Fatal(err)
	}
}
