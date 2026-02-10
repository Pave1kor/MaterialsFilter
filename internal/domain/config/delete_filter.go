package config

import (
	"fmt"
)

func deleteFilter(filters *[]Filter) error {
	fmt.Println()
	fmt.Println("Удаление фильтра по имени.")
	for {
		listFilters(*filters)
		fmt.Print("Введите имя фильтра, который необходимо удалить: ")
		name, err := newLine()
		if err != nil {
			return err
		}

		for i, filter := range *filters {
			if filter.Name == name {
				*filters = append((*filters)[:i], (*filters)[i+1:]...)
				fmt.Printf("Фильтр %s успешно удален!", name)
				return nil
			}
		}
		fmt.Println("Введенного имени фильтра не найдено, попробуйте еще раз!")
	}
}
