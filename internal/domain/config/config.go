package config

import (
	ptable "MaterialsFilter/internal/infrastructure/p_table"
	path "MaterialsFilter/internal/infrastructure/path"
	pathFile "MaterialsFilter/internal/infrastructure/path"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"os"
	"strings"
)

// Добавления элементов в существующий фильтр
func (cfg *Config) AddElementsInFilter() (string, error) {

	var newElement map[string]string

	//поиск фильтра
	nameFilter, err := findFilter(cfg.Filters)
	if err != nil {
		return "", err
	}

	//ввод списка элементов
	newElement, err = readElements()
	if err != nil {
		return "", err
	}

	//добавление новых элементов в фильтр
	for _, filter := range cfg.Filters {
		if filter.Name != nameFilter {
			continue
		}
		maps.Copy(filter.Filter, newElement)
		return filter.Name, nil
	}

	return "", nil

}

// Проверка: найти уникальный фильтр
func findFilter(filters []Filter) (string, error) {

	for {
		fmt.Print("Введите имя фильтра: ")
		nameFilter, err := NewLine()
		if err != nil {
			return "", err
		}

		if !existsFilter(nameFilter, filters) {
			fmt.Println("Введено неизвестное имя фильтра!")
			continue
		}

		fmt.Println("Желаете ли вы выбрать другой фильтр?")
		ok, err := Verification()
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

// Добавление нового фильтра
func (cfg *Config) AddNewFilter() (string, error) {
	newFilter, err := createNewFilter()
	if err != nil {
		return "", err
	}
	cfg.Filters = append(cfg.Filters, newFilter)
	return newFilter.Name, nil
}

// Создание нового фильтра
func createNewFilter() (Filter, error) {
	var filter Filter

	pathFile := path.New("..")
	output, err := pathFile.Output()
	if err != nil {
		return Filter{}, err
	}

	name, err := nameFilter()
	if err != nil {
		return Filter{}, err
	}

	filterLists, err := readElements()
	if err != nil {
		return Filter{}, err
	}

	filter.Filter = filterLists
	filter.Name = name
	filter.Output = output

	return filter, nil
}

// Ввод имени фильтра
func nameFilter() (string, error) {
	fmt.Print("Введите имя фильтра: ")
	name, err := NewLine()
	if err != nil {
		return "", err
	}
	return strings.Trim(name, " "), nil
}

// Проверка: существуют ли фильтры
func HasFilters(filters []Filter) bool {
	if len(filters) == 0 {
		fmt.Println()
		fmt.Println("Фильтры отсутвуют в файле настроек. Пожалуйста добавьте новый фильтр!")
		fmt.Println()
		return false
	}
	return true
}

// Проверка на наличие файла для обработки
func HasInputName(input string) bool {
	if input == "" {
		return false
	}
	return true
}

// Создание списка элементов для фильтрации
func readElements() (map[string]string, error) {
	listElements := make(map[string]string)
	fmt.Println()
	fmt.Println("Создание списка элементов для фильтрации.")
	for {
		fmt.Print("Введите элемент (или нажмите enter для завершения): ")

		element, err := NewLine()
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

// Изменение имени файла с входными данными
func (cfg *Config) ChangeInputFile() error {

	path := pathFile.New("..")
	inputPath, err := path.Input()
	if err != nil {
		return err
	}
	cfg.Input = inputPath
	return nil
}

// Удаление всех фильтров
func (cfg *Config) DeleteAllFilters() {
	cfg.Filters = nil
}

func (cfg *Config) DeleteElementsInFilter() error {

	var deleteElement string
	nameFilter, err := findFilter(cfg.Filters)
	if err != nil {
		return err
	}

	// удаление элементов из фильтра
	for i, filter := range cfg.Filters {
		if filter.Name == nameFilter {
			for {
				fmt.Print("Введите элемент (или нажмите enter для завершения): ")
				deleteElement, err = NewLine()
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
						delete((cfg.Filters)[i].Filter, deleteElement)
					}
				}
			}
		}
	}
	return nil
}

// Проверка существования фильтра по имени
func existsFilter(nameFilter string, filters []Filter) bool {
	for _, filter := range filters {
		if filter.Name == nameFilter {
			return true
		}
	}
	return false
}

func (cfg *Config) DeleteFilter() (string, error) {
	for {
		fmt.Print("Введите имя фильтра, который необходимо удалить: ")
		name, err := NewLine()
		if err != nil {
			return "", err
		}

		for i, filter := range cfg.Filters {
			if filter.Name == name {
				cfg.Filters = append((cfg.Filters)[:i], (cfg.Filters)[i+1:]...)

				return name, nil
			}
		}
		fmt.Println("Введенного имени фильтра не найдено, попробуйте еще раз!")
	}
}

func (cfg *Config) SaveConfigToJSON(configPath string) error {

	err := os.Remove(configPath)
	if err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(configPath, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Проверка: да/нет
func Verification() (bool, error) {
	for {
		fmt.Print("Ожидаю (да/нет) или (yes/no): ")

		input, err := NewLine()
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

// Ввод новой строки
func NewLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	newLine, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	newLine = strings.TrimSpace(newLine)
	return newLine, nil
}
