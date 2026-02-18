package config_manager

import (
	cfg "MaterialsFilter/internal/domain/config"
	cli "MaterialsFilter/internal/ui/cli"
	"fmt"
)

func ChangeConfig(config *cfg.Config) {
	if len(config.Input) == 0 {
		fmt.Println("Не задано имя файла с исходными данными.")
		cli.ChangeInputFileUI(config)
	}

	if len(config.Filters) == 0 {
		fmt.Println("Список фильтров пуст.")
		cli.AddNewFilterUI(config)
	}

	fmt.Println("Желаете ли вы изменить настройки фильтрации?")
	if ok, err := cli.Verification(); err != nil {
		fmt.Println("Ошибка при проверке ввода:", err)
		return
	} else if ok {
		changeFilter(config)
	}
}

func changeFilter(config *cfg.Config) {

	cli.CommandsInformation()
	for {
		fmt.Println()
		fmt.Print("Команда: ")
		command, err := cli.NewLine()
		if err != nil {
			fmt.Println("Ошибка при чтении ввода:", err)
			continue
		}

		switch command {
		case "addF":
			cli.AddNewFilterUI(config)
		case "delF":
			cli.DeleteFilterUI(config)
		case "delAllF":
			cli.DeleteAllFiltersUI(config)
		case "delEl":
			cli.DeleteElementsInFilterUI(config)
		case "addEl":
			cli.AddElementsInFilterUI(config)
		case "info":
			cli.InformationAboutConfig(*config)
		case "changeIn":
			cli.ChangeInputFileUI(config)
		case "command":
			cli.CommandsInformation()
		case "run":
			return
		default:
			fmt.Println("Введена неизвестная команда")
		}
	}
}
