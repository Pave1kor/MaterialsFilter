package materials

import (
	cfg "MaterialsFilter/config"
	filterObj "MaterialsFilter/internal/domain/filter"
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
func (m *MaterialsInformation) MaterialsFilter(filter cfg.Filter) {
	filterList := filterObj.FilterLists{}
	filterList.SetFilters(filter)

	for materialName, materialRow := range *m {
		materialRow.NewMaterialsRow()
		passed := true
		for _, elements := range materialRow.ParseMaterials {
			if !filterList.Get(filter.Name, elements) {
				passed = false
			}
			materialRow.FiltersName[filter.Name] = passed
		}
		(*m)[materialName] = materialRow
	}
}

// Сохранение отфильтрованного списка
func (m *MaterialsInformation) WriteCSV(filter cfg.Filter) error {
	file, err := os.Create(filter.Output)
	if err != nil {
		return err
	}
	defer file.Close()
	writeFile(file, filter.Name, m)
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
