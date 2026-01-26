package materials

import (
	cfg "MaterialsFilter/internal/config"
	filter "MaterialsFilter/internal/domain/filter"
	"encoding/csv"
	"io"
	"os"
	"regexp"
)

// Извлечение данных из файла
func (m *MaterialsInformation) ReaderCSV(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = ';'
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		(*m)[MaterialsName(record[0])] = MaterialRow{InformationElements: record}
	}
	return nil
}

// Парсинг химической формулы
func (m *MaterialsInformation) ParseMaterials() {
	for materialsName, material := range *m {
		material.ParseMaterials = regParsedMaterials(string(materialsName))
		(*m)[materialsName] = material
	}
}

var elementRegexp = regexp.MustCompile(`[A-Z][a-z]?`)

func regParsedMaterials(material string) []string {
	return elementRegexp.FindAllString(material, -1)
}

// фильтр соединений
func (m *MaterialsInformation) MaterialsFilter(filters []cfg.Filter) {
	filterList := filter.FilterLists{}
	filterList.SetFilters(filters)
	for materialName, materialRow := range *m {
		for _, filter := range filters {
			materialRow.NewMaterialsRow()
			passed := true
			for _, elements := range materialRow.ParseMaterials {
				if !filterList.Get(filter.NameFilter, elements) {
					passed = false
				}
				materialRow.FiltersName[filter.NameFilter] = passed
			}
			(*m)[materialName] = materialRow
		}
	}
}

// Сохранение отфильтрованного списка
func (m *MaterialsInformation) WriteCSV(filters []cfg.Filter) error {
	for _, filter := range filters {
		file, err := os.Create(filter.OutputData)
		if err != nil {
			return err
		}
		defer file.Close()
		writeFile(file, filter.NameFilter, m)
	}
	return nil
}

func writeFile(file *os.File, nameFilters string, m *MaterialsInformation) error {
	writer := csv.NewWriter(file)
	writer.Comma = ';'
	for _, material := range *m {
		if material.FiltersName[nameFilters] {
			writer.Write(material.InformationElements)
		}
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		return err
	}

	return nil
}
