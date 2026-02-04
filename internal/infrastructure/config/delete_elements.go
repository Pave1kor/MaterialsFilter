package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func deleteElementsInFilter(config *Config) error {

	if config.Filters == nil {
		fmt.Println("Фильтры отсутвуют в файле настроек. Пожалуйста добавьте новый фильтр!")
		return nil
	}

	var deleteElement string
	nameFilter, err := findFilter(config)
	if err != nil {
		return err
	}

	// удаление элементов из фильтра
	for i, filter := range config.Filters {
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
				if ok, _ := get(deleteElement); !ok {
					fmt.Println("Ввведен неизвестный элемент, попробуйте еще раз.")
					continue
				} else {
					if _, exists := filter.Filter[deleteElement]; !exists {
						fmt.Println("Введенный элемент отсутствует в фильтре, попробуйте еще раз.")
						continue
					} else {
						delete(config.Filters[i].Filter, deleteElement)
					}
				}
			}
			fmt.Println("Элементы успешно удалены из фильтра")
			listElements(config, nameFilter)
		}
	}
	return nil
}

// проверка существования фильтра по имени
func existsFilter(nameFilter string, config *Config) bool {
	for _, filter := range config.Filters {
		if filter.Name == nameFilter {
			return true
		}
	}
	return false
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
