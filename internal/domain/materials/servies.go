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
func (m *MaterialsInformation) MaterialsFilter(config cfg.Config) {
	filterList := filter.FilterLists{}
	filterList.SetFilters(config)
	for materialName, materialRow := range *m {
		for nameFilter := range config.Filters {
			materialRow.NewMaterialsRow()
			passed := true
			for _, elements := range materialRow.ParseMaterials {
				if !filterList.Get(nameFilter, elements) {
					passed = false
				}
				materialRow.FiltersName[nameFilter] = passed
			}
			(*m)[materialName] = materialRow
		}
	}
}

// Сохранение отфильтрованного списка
func (m *MaterialsInformation) WriteCSV(config cfg.Config) error {
	for nameFilters := range config.Filters {
		path := config.Paths.OutputData + nameFilters + ".csv"
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
		writeFile(file, nameFilters, m)
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
