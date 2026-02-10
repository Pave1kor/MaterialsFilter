package config

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// проверка: да/нет
func verification() (bool, error) {
	for {
		fmt.Print("Ожидаю (да/нет) или (yes/no): ")

		input, err := newLine()
		if err != nil {
			return false, err
		}

		switch input {
		case "да", "yes", "y":
			return true, nil
		case "нет", "no", "n":
			return false, nil
		case "":
			return false, errors.New("ввод отменён")
		default:
			fmt.Println("Введите 'да' или 'нет'.")
		}
	}
}

// проверка: существуют ли фильтры
func hasFilters(filters []Filter) bool {
	if len(filters) == 0 {
		fmt.Println()
		fmt.Println("Фильтры отсутвуют в файле настроек. Пожалуйста добавьте новый фильтр!")
		fmt.Println()
		return false
	}
	return true
}

// чтение новой строки из терминала
func newLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	newLine, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	newLine = strings.TrimSpace(newLine)
	return newLine, nil
}

// вывод списка фильтров
func listFilters(filters []Filter) {
	fmt.Println()
	fmt.Println("Список фильтров следующий:")
	var filterList []string

	for _, nameF := range filters {
		filterList = append(filterList, nameF.Name)
	}
	fmt.Println(strings.Join(filterList, ", "))
	fmt.Println()
}
