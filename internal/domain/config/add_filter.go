package config

import (
	ptable "MaterialsFilter/internal/infrastructure/p_table"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// добавление нового фильтра
func addNewFilter(config *Config) error {

	newFiter, err := createNewFilter(*config)
	if err != nil {
		return err
	}
	config.Filters = append(config.Filters, newFiter)
	fmt.Printf("Фильтр %s успешно создан!\n", newFiter.Name)
	listFilters(*config)
	return nil
}

// проверяется уникальность введенного имени фильтра
func findUnicName(config Config, name string) bool {
	if len(config.Filters) == 0 {
		return true
	}
	for _, filter := range config.Filters {
		if filter.Name == name {
			fmt.Println("Название фильта должно быть уникальным!")
			listFilters(config)
			return false
		}
	}
	return true
}

func listFilters(config Config) {
	fmt.Println("Список фильтров следующий:")
	var filterList []string

	for _, nameF := range config.Filters {
		filterList = append(filterList, nameF.Name)
	}
	fmt.Println(strings.Join(filterList, ", "))
}

// создается новый фильтр
func createNewFilter(config Config) (Filter, error) {
	var filter Filter

	name, err := createNameFilter(config)
	if err != nil {
		return Filter{}, err
	}

	output, err := filepath.Abs(filepath.Clean(filepath.Join("..", "data", "output", name+".csv")))
	if err != nil {
		return Filter{}, err
	}

	filterLists, err := addElements()
	if err != nil {
		return Filter{}, err
	}

	filter.Filter = filterLists
	filter.Name = name
	filter.Output = output

	return filter, nil
}

// создание списка элементов для фильтрации
func addElements() (map[string]string, error) {
	listElements := make(map[string]string)

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

// Добавление фильтра с уникальным именем
func createNameFilter(config Config) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println()
		fmt.Print("Введите имя фильтра: ")
		name, err := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		if err != nil {
			return "", err
		}
		if findUnicName(config, name) {
			return name, nil
		}
	}
}
