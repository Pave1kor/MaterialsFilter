package config

import "fmt"

func deleteAllFilters(filters *[]Filter) {
	fmt.Println()
	fmt.Println("Удаление всех фильтров.")
	*filters = nil
}
