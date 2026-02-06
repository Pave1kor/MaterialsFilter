package csvfile

import (
	"encoding/csv"
	"os"
)

// в дальнейшем можно вынести в infrastructure. Здесь создаем интерфейс

// Сохранение отфильтрованного списка
func (obj *CSVFile) WriteCSV(filteredData map[string][]string, filterName string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	writeFile(file, filteredData, createFlterName(filterName))
	return nil
}

func createFlterName(filterName string) []string {
	nameArr := [...]string{"Имя фильтра", filterName}
	return nameArr[:]
}

func writeFile(file *os.File, filteredData map[string][]string, filterName []string) error {
	writer := csv.NewWriter(file)
	writer.Comma = ';'

	writer.Write(filterName)
	for _, information := range filteredData {
		writer.Write(information)
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return err
	}

	return nil
}
