package filter

import (
	regexp "MaterialsFilter/internal/infrastructure/regexp"
)

// Фильтр соединений
func ElementsFilter(data [][]string, filter []string) [][]string {
	results := [][]string{}
	for _, information := range data {
		elements, err := regexp.RegexpFilter(information[0])
		if err != nil {
			continue
		}
		if containOnlyAllowed(elements, filter) {
			results = append(results, information)
		}
	}
	return results
}

// Обработка данных по соответствующему фильтру
func containOnlyAllowed(elements []string, filter []string) bool {
	if len(elements) == 0 {
		return false
	}
	filterMap := make(map[string]struct{})
	for _, element := range filter {
		filterMap[element] = struct{}{}
	}

	for _, element := range elements {
		if _, found := filterMap[element]; !found {
			return false
		}
	}
	return true
}
