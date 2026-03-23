package app

import (
	cfg "MaterialsFilter/internal/domain/config"
	filters "MaterialsFilter/internal/domain/filter"
	readerWriterCSV "MaterialsFilter/internal/infrastructure/csv"
	cli "MaterialsFilter/internal/ui/cli"
	"fmt"
)

func Run(config *cfg.Config) {

	// Загрузка  (разделитель и имя файла для обработки)
	csvFile := readerWriterCSV.NewCSVFile(config.InputPathFile)

	// Чтение csv файла
	csvFile.ReadCSV()

	// Вывод таблицы с данными (шапка и одна строка) в теримнал
	cli.ViewTable(csvFile.Headlines, csvFile.Table)

	fmt.Println("Желаете ли вы изменить заголовки столбцов?")
	if cli.Verification() {
		cli.ChangeHeadlinesUI(csvFile)
	}

	fmt.Println("Выполняется фильтрация по заданным фильтрам.")
	// Фильтрация данных в соответствии с заданными настройками
	for _, filter := range config.Filters {
		filteredData := filters.ElementsFilter(csvFile.Data, filter.Elements)
		csvFile.WriteCSV(filteredData, filter.Elements, filter.Name, filter.OutputPathFile, csvFile.Headlines)
	}
	fmt.Println("Успех! Нажмите клавишу `Enter`, чтобы закрыть программу.")
	cli.NewLine()
}
