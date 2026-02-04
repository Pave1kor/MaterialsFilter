package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func changeConfig(config *Config) error {
	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Print("\nВведите команду: ")
		command, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		command = strings.TrimSpace(command)
		switch command {
		case "addF":
			addNewFilter(config) // добавить новый фильтр
		case "delF":
			deleteFilter(config) // удалить фильтр по названию
		case "delAllF":
			deleteAllFilters(config) // удалить все фильтры
		case "delEl":
			deleteElementsInFilter(config) // удалить желаемые элементы из списка
		case "addEl":
			addElementsInFilter(config) //добавить элементы в существующий фильтр
		case "info":
			informationAboutConfig(config) // вывести информацию обо всех фильтрах
		case "command":
			commandsInformation()
		case "run":
			if config.Filters == nil {
				fmt.Println("Фильтры отсутвуют в файле настроек. Пожалуйста добавьте новый фильтр!")
				continue
			}
			MarshalingConfig(*config)
			return nil
		default:
			fmt.Println("Введена неизвестная команда")
		}
	}
}
