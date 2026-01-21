package materials

import (
	filter "MaterialsFilter/internal/filter"
	"encoding/csv"
	"io"
	"os"
	"regexp"
	"slices"
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
		(*m)[MaterialsName(record[0])] = Material{InformationElements: record}
	}
	return nil
}

// парсинг химической формулы
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
func (m *MaterialsInformation) MaterialsFilter() {
	m.filter(*filter.NewHeuslerFilter(), func(m *Material, v bool) {
		m.FilterHeusler = v
	})
	m.filter(*filter.NewChalcogenideFilter(), func(m *Material, v bool) {
		m.FilterChalcogenide = v
	})
}

func (m *MaterialsInformation) filter(
	filterList filter.FilterLists,
	setter func(*Material, bool),
) {
	for name, material := range *m {
		ok := !slices.ContainsFunc(material.ParseMaterials, filterList.Get)
		setter(&material, ok)
		(*m)[name] = material
	}
}

// Сохранение отфильтрованного списка Халькогенидов
func (m *MaterialsInformation) WriteCSVChalcogenide(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = ';'
	for _, material := range *m {
		if material.FilterChalcogenide {
			writer.Write(material.InformationElements)
		}
	}
	writer.Flush()
	if err = writer.Error(); err != nil {
		return err
	}

	return nil
}

// Сохранение отфильтрованного списка сплавов Гейслера
func (m *MaterialsInformation) WriteCSVHeusler(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = ';'
	for _, material := range *m {
		if material.FilterHeusler {
			writer.Write(material.InformationElements)
		}
	}
	writer.Flush()
	if err = writer.Error(); err != nil {
		return err
	}
	return nil
}
