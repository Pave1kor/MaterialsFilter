package csv

import (
	"encoding/csv"
	"os"
	"slices"
	"strings"
)

// Сохранение отфильтрованного списка
func (obj *CSVFile) WriteCSV(filteredData map[string][]string, listElements map[string]struct{}, filterName string, path string, headlines []string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writeFile(file, createFlterName(filterName), createListElements(listElements), filteredData, headlines)
	return nil
}

func createListElements(listElements map[string]struct{}) []string {
	keys := make([]string, 0, len(listElements))
	for k := range listElements {
		keys = append(keys, k)
	}

	slices.Sort(keys)

	return []string{
		"Список элементов для фильтрации",
		strings.Join(keys, ", "),
	}
}

func createFlterName(filterName string) []string {
	return []string{
		"Имя фильтра",
		filterName,
	}
}

func writeFile(file *os.File, filterName []string, listElements []string, filteredData map[string][]string, headlines []string) error {
	writer := csv.NewWriter(file)
	writer.Comma = ';'

	writer.Write(filterName)
	writer.Write(listElements)
	writer.Write(headlines)
	for _, information := range filteredData {
		writer.Write(information)
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return err
	}

	return nil
}
