package filter

import (
	config "MaterialsFilter/internal/config"
	"regexp"
)

// Загрузка фильтров из конфига
func (s *FilterLists) SetFilters(filterLists []config.Filter) {
	s.NewFilter()
	re := regexp.MustCompile(`[A-Z][a-z]?`)

	for _, filter := range filterLists {
		valuesListsArr := re.FindAllString(filter.DataFilter, -1)
		var valuesListMap = make(map[Elements]any, 100)
		for _, valuesLists := range valuesListsArr {
			valuesListMap[Elements(valuesLists)] = struct{}{}
		}
		s.listFilters[filter.NameFilter] = valuesListMap
	}
}

// Фильтрация значения
func (s *FilterLists) Get(nameFilter, material string) bool {
	_, ok := s.listFilters[nameFilter][Elements(material)]
	return ok
}

func (s *FilterLists) NewFilter() {
	if s.listFilters == nil {
		s.listFilters = make(map[string]map[Elements]any)
	}
}
