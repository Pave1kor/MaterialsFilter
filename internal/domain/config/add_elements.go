package config

import (
	"errors"
	"fmt"
	"maps"
)

// добавления элементов в существующий фильтр
func addElementsInFilter(filters *[]Filter) error {
	fmt.Println()
	fmt.Println("Добавление новых элементов в фильтр.")
	var newElement map[string]string

	//поиск фильтра
	nameFilter, err := findFilter(*filters)
	if err != nil {
		return err
	}

	//ввод списка элементов
	newElement, err = readElements()
	if err != nil {
		return err
	}

	//добавление новых элементов в фильтр
	for _, filter := range *filters {
		if filter.Name != nameFilter {
			continue
		}
		maps.Copy(filter.Filter, newElement)
		fmt.Printf("Введенные элементы успешно добавлены в фильтр %s.\n", filter.Name)
		break
	}

	return nil

}

// проверка: найти уникальный фильтр
func findFilter(filters []Filter) (string, error) {
	listFilters(filters)
	for {

		fmt.Print("Введите имя фильтра: ")
		nameFilter, err := newLine()
		if err != nil {
			return "", err
		}

		if !existsFilter(nameFilter, filters) {
			fmt.Println("Введено неизвестное имя фильтра!")
			continue
		}

		listElements(filters, nameFilter)
		fmt.Println("Желаете ли вы выбрать другой фильтр?")
		ok, err := verification()
		if errors.Is(err, errors.New("ввод отменён")) {
			return "", err
		}
		if err != nil {
			return "", err
		}
		if !ok {
			return nameFilter, nil
		}
	}
}
