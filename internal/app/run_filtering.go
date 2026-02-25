package app

import (
	cfg "MaterialsFilter/internal/domain/config"
	filters "MaterialsFilter/internal/domain/filter"
	readerWriterCSV "MaterialsFilter/internal/infrastructure/csv"
	cli "MaterialsFilter/internal/ui/cli"
	"fmt"
)

func Run(config *cfg.Config) error {

	// Создание объекта csv с начальными настройками (разделитель и имя файла для обработки)
	csvFile := readerWriterCSV.NewCSVFile(config.Input)

	// Получение данных из csv файла
	data, err := csvFile.ReadCSV()
	if err != nil {
		return err
	}

	// Получение таблицы с данными (шабка и одна строка)
	cli.ViewTable(csvFile.Headlines, csvFile.Table)

	fmt.Println("Желаете ли вы изменить заголовки столбцов?")
	confirmed, err := cli.Verification()
	if confirmed {
		cli.ChangeHeadlinesUI(csvFile)
	}

	// Фильтрация данных в соответствии с заданными настройками
	for _, filter := range config.Filters {
		filteredData := filters.ElementsFilter(data, filter.Filter)
		if err := csvFile.WriteCSV(filteredData, filter.Filter, filter.Name, filter.Output, csvFile.Headlines); err != nil {
			return err
		}
	}
	fmt.Println("Успех!")
	return nil
}
