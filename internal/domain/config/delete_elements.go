package config

import (
	ptable "MaterialsFilter/internal/infrastructure/p_table"
	"fmt"
)

func deleteElementsInFilter(filters *[]Filter) error {
	fmt.Println()
	fmt.Println("Удаление элементов из фильтра.")
	listFilters(*filters)
	var deleteElement string
	nameFilter, err := findFilter(*filters)
	if err != nil {
		return err
	}

	// удаление элементов из фильтра
	for i, filter := range *filters {
		if filter.Name == nameFilter {
			for {
				fmt.Print("Введите элемент (или нажмите enter для завершения): ")
				deleteElement, err = newLine()
				if err != nil {
					return err
				}
				if deleteElement == "" {
					break
				}
				if ok, _ := ptable.Get(deleteElement); !ok {
					fmt.Println("Ввведен неизвестный элемент, попробуйте еще раз.")
					continue
				} else {
					if _, exists := filter.Filter[deleteElement]; !exists {
						fmt.Println("Введенный элемент отсутствует в фильтре, попробуйте еще раз.")
						continue
					} else {
						delete((*filters)[i].Filter, deleteElement)
					}
				}
			}
			fmt.Println("Элементы успешно удалены из фильтра")
			listElements(*filters, nameFilter)
		}
	}
	return nil
}

// проверка существования фильтра по имени
func existsFilter(nameFilter string, filters []Filter) bool {
	for _, filter := range filters {
		if filter.Name == nameFilter {
			return true
		}
	}
	return false
}
