package csvfile

import (
	"encoding/csv"
	"os"
)

// в дальнейшем можно вынести в infrastructure. Здесь создаем интерфейс

// Сохранение отфильтрованного списка
func (obj *CSVFile) WriteCSV(filteredData map[string][]string, listElements map[string]string, filterName string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	writeFile(file, createFlterName(filterName), createListElements(listElements), filteredData)
	return nil
}

func createListElements(listElements map[string]string) []string {
	var list = make([]string, 0, len(listElements))
	list = append(list, "Список элементов для фильтрации")
	for element := range listElements {
		list = append(list, element)
	}
	return list
}

func createFlterName(filterName string) []string {
	nameArr := [...]string{"Имя фильтра", filterName}
	return nameArr[:]
}

func writeFile(file *os.File, filterName []string, listElements []string, filteredData map[string][]string) error {
	writer := csv.NewWriter(file)
	writer.Comma = ';'
	writer.Write(filterName)
	writer.Write(listElements)

	for _, information := range filteredData {
		writer.Write(information)
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return err
	}

	return nil
}
