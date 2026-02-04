package materials

import (
	regexp "MaterialsFilter/internal/infrastructure/filter"
)

// фильтр соединений
func ElementsFilter(data map[string][]string, filter map[string]string) map[string][]string {
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

// обработка данных по соответствующему фильтру

func containOnlyAllowed(elements []string, filter map[string]string) bool {
	if len(elements) == 0 {
		return false
	}
	for _, element := range elements {
		if _, exists := filter[element]; !exists {
			return false
		}
	}
	return true
}
