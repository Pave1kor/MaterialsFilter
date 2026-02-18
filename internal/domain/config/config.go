package config

import (
	"fmt"
)

// Добавления элементов в существующий фильтр
func (cfg *Config) AddElementsInFilter(nameFilter string, listElements []string) {
	for _, filter := range cfg.Filters {
		if filter.Name != nameFilter {
			continue
		}
		for _, element := range listElements {
			filter.Filter[element] = struct{}{}
		}
	}
}

// Добавление нового фильтра
func (cfg *Config) AddNewFilter(output string, nameFilter string, listElements []string) error {

	newFilter, err := createNewFilter(output, nameFilter, listElements)
	if err != nil {
		return err
	}
	cfg.Filters = append(cfg.Filters, newFilter)
	return nil
}

// Создание нового фильтра
func createNewFilter(output string, nameFilter string, listElements []string) (Filter, error) {
	var filter Filter
	mapElements := make(map[string]struct{}, 1000)

	for _, element := range listElements {
		mapElements[element] = struct{}{}
	}
	filter.Filter = mapElements
	filter.Name = nameFilter
	filter.Output = output

	return filter, nil
}

// Проверка: существуют ли фильтры
func HasFilters(filters []Filter) bool {
	return len(filters) > 0
}

// Проверка на наличие файла для обработки
func HasInputName(input string) bool {
	return len(input) > 0
}

// Изменение имени файла с входными данными
func (cfg *Config) ChangeInputFile(inputPath string) {
	cfg.Input = inputPath
}

// Удаление всех фильтров
func (cfg *Config) DeleteAllFilters() {
	cfg.Filters = nil
}

// удаление элементов из фильтра
func (cfg *Config) DeleteElementsInFilter(nameFilter string, listElements []string) {

	for i, filter := range cfg.Filters {
		if filter.Name != nameFilter {
			continue
		}
		for _, element := range listElements {
			delete((cfg.Filters)[i].Filter, element)
		}
	}
}

// Проверка существования фильтра по имени
func NewFilter(nameFilter string, filters []Filter) bool {
	for _, filter := range filters {
		if filter.Name == nameFilter {
			return false
		}
	}
	return true
}
func (cfg *Config) DeleteFilter(name string) error {
	for i, filter := range cfg.Filters {
		if filter.Name == name {
			cfg.Filters = append((cfg.Filters)[:i], (cfg.Filters)[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Введенного имени фильтра не найдено, попробуйте еще раз!")
}

// Проверка существования фильтра по имени
func (cfg Config) ExistFilter(nameFilter string) bool {
	for _, filter := range cfg.Filters {
		if filter.Name == nameFilter {
			return true
		}
	}
	return false
}
