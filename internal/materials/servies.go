package materials

import (
	filter "MaterialsFilter/internal/filter"
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"regexp"
	"strings"
)

// Извлечение данных из файла
func (m *MaterialsInformation) ReaderCSV() error{
	file, err := os.Open("ZT value.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		row := strings.Split(record[0], ";")
		(*m)[MaterialsName(row[0])] = Material{InformationElements: strings.Join(record, "")}
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
func (m *MaterialsInformation) MaterialsFilterHeusler() {
	filterLists := filter.HeuslerFilter()
	for materialsName, material := range *m {
		for _, element := range material.ParseMaterials {
			material.FilterHeusler = true
			if _, ok := filterLists[filter.Elements(element)]; !ok {
				material.FilterHeusler = false
				break
			}
		}
		(*m)[materialsName] = material
	}
}
func (m *MaterialsInformation) MaterialsFilterChalcogenide() {
	filterLists := filter.ChalcogenideFilter()
	for materialsName, material := range *m {
		for _, element := range material.ParseMaterials {
			material.FilterChalcogenide = true
			if _, ok := filterLists[filter.Elements(element)]; !ok {
				material.FilterChalcogenide = false
				break
			}
		}
		(*m)[materialsName] = material
	}
}

// сохранение отфильтрованных значений
func (m *MaterialsInformation) WriteCSVChalcogenide() error{
	file, err := os.Create("filteredMaterialsChalcogenide.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, material := range *m {
		if material.FilterChalcogenide {
			writer.WriteString(material.InformationElements)
			writer.WriteByte('\n')
		}
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
return nil
}

func (m *MaterialsInformation) WriteCSVHeusler() error {
	file, err := os.Create("filteredMaterialsHeusler.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, material := range *m {
		if material.FilterHeusler {
			writer.WriteString(material.InformationElements)
			writer.WriteByte('\n')
		}
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}
