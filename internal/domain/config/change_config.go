package config

import (
	"fmt"
)

func ChangeConfig(config *Config) {
	if !hasFilters(config.Filters) {
		addNewFilter(config.Filters)
	}

	fmt.Println("Желаете ли вы изменить настройки фильтрации?")
	if ok, err := verification(); err != nil {
		fmt.Println("Ошибка при проверке ввода:", err)
		return
	} else if ok {
		changeFilter(config)
	}
}

func changeFilter(config *Config) {

	commandsInformation()
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
			addNewFilter(config.Filters)
		case "delF":
			deleteFilter(&config.Filters)
		case "delAllF":
			deleteAllFilters(&config.Filters)
		case "delEl":
			deleteElementsInFilter(&config.Filters)
		case "addEl":
			addElementsInFilter(&config.Filters)
		case "info":
			informationAboutConfig(config)
		case "changeIn":
			changeInputFile(config)
		case "command":
			commandsInformation()
		case "run":
			return
		default:
			fmt.Println("Введена неизвестная команда")
		}
	}
}
