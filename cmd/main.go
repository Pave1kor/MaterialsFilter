package main

import (
	materials "MaterialsFilter/internal/materials"
	"log"
)

func main() {
	material := materials.NewMaterials()
	if err := material.ReaderCSV("../data/input/zt_value.csv"); err != nil {
		log.Fatal(err)
	}
	material.ParseMaterials()
	material.MaterialsFilter()
	if err := material.WriteCSVHeusler("../data/output/HeuslerFilter.csv"); err != nil {
		log.Fatal(err)
	}
	if err := material.WriteCSVChalcogenide("../data/output/ChalcogenideFilter.csv"); err != nil {
		log.Fatal(err)
	}
}
