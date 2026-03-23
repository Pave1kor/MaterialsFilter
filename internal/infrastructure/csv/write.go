package csv

import (
	"encoding/csv"
	"log"
	"os"
	"slices"
	"strings"
)

// Сохранение отфильтрованного списка
func (obj *CSVFile) WriteCSV(filteredData [][]string, listElements []string, filterName string, path string, headlines []string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writeFile(file, createFlterName(filterName), createListElements(listElements), filteredData, headlines)
}

func createListElements(listElements []string) []string {
	slices.Sort(listElements)
	return []string{
		"List of elements to filter",
		strings.Join(listElements, ", "),
	}
}

func createFlterName(filterName string) []string {
	return []string{
		"Filter name",
		filterName,
	}
}

func writeFile(file *os.File, filterName []string, listElements []string, filteredData [][]string, headlines []string) {
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
		log.Fatal(err)
	}
}
