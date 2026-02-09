package config

import "fmt"

func deleteAllFilters(config *Config) {
	if len(config.Filters) == 0 {
		fmt.Println("Фильтры отсутвуют в файле настроек. Пожалуйста добавьте новый фильтр!")
		return
	}
	config.Filters = nil
}
