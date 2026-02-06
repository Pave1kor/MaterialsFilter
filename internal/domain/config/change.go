package config

import (
	"fmt"
)

func ChangeConfig(config *Config) {
	fmt.Println("Желаете ли вы изменить настройки фильтрации? Введите 'да' или 'нет'.")
	if ok, err := Verification(); err != nil {
		fmt.Println("Ошибка при проверке ввода:", err)
		return
	} else if ok {
		changeFilter(config)
	}
}

func changeFilter(config *Config) {
	fmt.Println("Введите команду для изменения фильтров:")
	CommandsInformation()
	for {
		fmt.Println()
		fmt.Print("Команда: ")
		command, err := newLine()
		if err != nil {
			fmt.Println("Ошибка при чтении ввода:", err)
			continue
		}
		switch command {
		case "addF":
			addNewFilter(config)
		case "delF":
			deleteFilter(config)
		case "delAllF":
			deleteAllFilters(config)
		case "delEl":
			deleteElementsInFilter(config)
		case "addEl":
			addElementsInFilter(config)
		case "info":
			informationAboutConfig(config)
		case "command":
			CommandsInformation()
		case "run":
			if config.Filters == nil {
				fmt.Println("Фильтры отсутвуют в файле настроек. Пожалуйста добавьте новый фильтр!")
				continue
			}
			return
		default:
			fmt.Println("Введена неизвестная команда")
		}
	}
}
