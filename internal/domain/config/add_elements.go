package config

import (
	"fmt"
	"maps"
)

// добавления элементов в существующий фильтр
func addElementsInFilter(config *Config) error {
	var newElement map[string]string

	if config.Filters == nil {
		fmt.Println("Фильтры отсутвуют в файле настроек. Пожалуйста добавьте новый фильтр!")
		return nil
	}

	//поиск фильтра
	nameFilter, err := findFilter(config)
	if err != nil {
		return err
	}

	//ввод списка элементов
	newElement, err = addElements()
	if err != nil {
		return err
	}

	//добавление новых элементов в фильтр
	for _, filter := range config.Filters {
		if filter.Name == nameFilter {
			maps.Copy(filter.Filter, newElement)
			fmt.Printf("Введенные элементы успешно добавлены в фильтр %s.\n", filter.Name)
		}
	}

	return nil

}

func findFilter(config *Config) (string, error) {
	for {
		fmt.Print("Введите имя фильтра: ")
		nameFilter, err := newLine()
		if err != nil {
			return "", err
		}
		if !existsFilter(nameFilter, config) {
			fmt.Println("Введено неизвестное имя фильтра!")
		} else {

			listElements(config, nameFilter)
			fmt.Println("Желаете ли вы изменить имя фильтра? Введите 'да' или 'нет'.")
			ok, err := Verification()
			if err != nil {
				return "", err
			}
			if !ok {
				return nameFilter, nil
			}
		}
	}
}

func Verification() (bool, error) {
	for {
		fmt.Print("Ожидаю: ")
		ok, err := newLine()
		if err != nil {
			return false, err
		}
		switch ok {
		case "да":
			return true, nil
		case "нет":
			return false, nil
		default:
			fmt.Println("Некорректный ввод. Пожалуйста, введите 'да' или 'нет'.")
			continue
		}
	}
}
