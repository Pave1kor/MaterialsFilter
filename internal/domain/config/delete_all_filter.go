package config

import "fmt"

func deleteAllFilters(config *Config) {
	if config.Filters == nil {
		fmt.Println("Фильтры отсутвуют в файле настроек. Пожалуйста добавьте новый фильтр!")
		return
	}
	config.Filters = nil
}
