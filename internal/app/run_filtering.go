package app

import (
	readerWriterCSV "MaterialsFilter/internal/domain/csv_file"
	filters "MaterialsFilter/internal/domain/filter"
	cfg "MaterialsFilter/internal/infrastructure/config"
	"log"
)

func Run(config *cfg.Config) {
	for _, file := range config.Filters {
		csvFile := readerWriterCSV.NewCSVFile(file.Input, file.Output)
		data, err := csvFile.ReadCSV()
		if err != nil {
			log.Fatal(err)
		}

		filteredData := filters.ElementsFilter(data, file.Filter)
		if err := csvFile.WriteCSV(filteredData); err != nil {
			log.Fatal(err)
		}
	}
}
