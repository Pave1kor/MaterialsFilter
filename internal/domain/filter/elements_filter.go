package materials

import (
	regexp "MaterialsFilter/internal/infrastructure/filter"
)

// фильтр соединений
func ElementsFilter(data map[string][]string, filter map[string]struct{}) map[string][]string {
	results := make(map[string][]string)
	for formula, information := range data {
		elements, err := regexp.RegexpFilter(formula)
		if err != nil {
			continue
		}
		if containOnlyAllowed(elements, filter) {
			results[formula] = information
		}
	}
	return results
}

func containOnlyAllowed(elements []string, filter map[string]struct{}) bool {
	for _, element := range elements {
		if _, exists := filter[element]; !exists {
			return false
		}
	}
	return true
}
