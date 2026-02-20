package app

import (
	cfg "MaterialsFilter/internal/domain/config"
	filters "MaterialsFilter/internal/domain/filter"
	readerWriterCSV "MaterialsFilter/internal/infrastructure/csv"
	cli "MaterialsFilter/internal/ui/cli"
	"fmt"
)

func Run(config *cfg.Config) error {
	csvFile := readerWriterCSV.NewCSVFile(config.Input)
	data, err := csvFile.ReadCSV()
	if err != nil {
		return err
	}

	fmt.Println("Желаете ли вы изменить заголовки столбцов?")
	confirmed, err := cli.Verification()
	if confirmed {
		cli.ChangeHeadlinesUI(csvFile)
	}
	for _, filter := range config.Filters {
		filteredData := filters.ElementsFilter(data, filter.Filter)
		if err := csvFile.WriteCSV(filteredData, filter.Filter, filter.Name, filter.Output, csvFile.Headlines); err != nil {
			return err
		}
	}
	fmt.Println("Успех!")
	return nil
}
