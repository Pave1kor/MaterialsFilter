package config

import (
	ptable "MaterialsFilter/internal/infrastructure/p_table"
	path "MaterialsFilter/internal/infrastructure/path"
	"fmt"
	"path/filepath"
	"strings"
)

// добавление нового фильтра
func addNewFilter(filters []Filter) error {
	fmt.Println("Добавление нового фильтра.")
	newFilter, err := createNewFilter()
	if err != nil {
		return err
	}
	filters = append(filters, newFilter)
	fmt.Printf("Фильтр %s успешно создан!\n", newFilter.Name)
	listFilters(filters)
	return nil
}

// создается новый фильтр
func createNewFilter() (Filter, error) {
	var filter Filter

	pathFile := path.New("..")
	output, err := pathFile.Output()
	if err != nil {
		return Filter{}, err
	}

	filterLists, err := readElements()
	if err != nil {
		return Filter{}, err
	}

	filter.Filter = filterLists
	filter.Name = strings.TrimRight(filepath.Base(output), ".csv")
	filter.Output = output

	return filter, nil
}

// создание списка элементов для фильтрации
func readElements() (map[string]string, error) {
	listElements := make(map[string]string)
	fmt.Println()
	fmt.Println("Создание списка элементов для фильтрации.")
	for {
		fmt.Print("Введите элемент (или нажмите enter для завершения): ")

		element, err := newLine()
		if err != nil {
			return nil, err
		}

		if element == "" {
			return listElements, nil
		}

		if ok, name := ptable.Get(element); !ok {
			fmt.Printf("Элемент %s не найден в периодической таблице Менделеева. Попробуйте снова.\n", element)
			continue
		} else {
			if _, ok := listElements[element]; ok {
				fmt.Println("Введенный элемент уже существует в фильтре, попробуйте еще раз!")
				continue
			}
			listElements[element] = name
		}
	}
}
