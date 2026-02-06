package app

import (
	cfg "MaterialsFilter/internal/domain/config"
	readerWriterCSV "MaterialsFilter/internal/domain/csv_file"
	filters "MaterialsFilter/internal/domain/filter"
	"log"
)

func Run(config *cfg.Config) {
	csvFile := readerWriterCSV.NewCSVFile(config.Input)
	data, err := csvFile.ReadCSV()
	if err != nil {
		log.Fatal(err)
	}
	for _, filter := range config.Filters {
		filteredData := filters.ElementsFilter(data, filter.Filter)
		if err := csvFile.WriteCSV(filteredData, filter.Name, filter.Output); err != nil {
			log.Fatal(err)
		}
	}

}
