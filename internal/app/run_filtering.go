package app

import (
	cfg "MaterialsFilter/internal/domain/config"
	filters "MaterialsFilter/internal/domain/filter"
	readerWriterCSV "MaterialsFilter/internal/infrastructure/csv"
)

func Run(config *cfg.Config) error {
	csvFile := readerWriterCSV.NewCSVFile(config.Input)
	data, err := csvFile.ReadCSV()
	if err != nil {
		return err
	}
	for _, filter := range config.Filters {
		filteredData := filters.ElementsFilter(data, filter.Filter)
		if err := csvFile.WriteCSV(filteredData, filter.Filter, filter.Name, filter.Output); err != nil {
			return err
		}
	}
	return nil
}
