package csvfile

import (
	"encoding/csv"
	"os"
)

// в дальнейшем можно вынести в infrastructure.Сдесь создаем интерфейс

// Сохранение отфильтрованного списка
func (obj *CSVFile) WriteCSV(filteredData map[string][]string) error {
	file, err := os.Create(obj.Output)
	if err != nil {
		return err
	}
	defer file.Close()

	writeFile(file, filteredData)
	return nil
}

func writeFile(file *os.File, filteredData map[string][]string) error {
	writer := csv.NewWriter(file)
	writer.Comma = ';'
	for _, information := range filteredData {
		writer.Write(information)
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return err
	}

	return nil
}
