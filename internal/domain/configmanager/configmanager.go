package config_manager

import (
	cfg "MaterialsFilter/internal/domain/config"
	cli "MaterialsFilter/internal/ui/cli"
	"fmt"
)

func ChangeConfig(config *cfg.Config) {

	if len(config.InputPathFile) == 0 {
		fmt.Println("Не задано имя файла с исходными данными.")
		cli.ChangeInputFileUI(config)
	}

	if len(config.Filters) == 0 {
		cli.AddNewFilterUI(config)
	}

	fmt.Println("Желаете ли вы изменить настройки фильтрации?")
	if cli.Verification() {
		changeFilter(config)
	}
	fmt.Println("Окей, переходим к фильтрации.")
}

func changeFilter(config *cfg.Config) {

	cli.CommandsInformation()
	for {
		fmt.Println()
		fmt.Print("Команда: ")
		command := cli.NewLine()

		switch command {
		case "add-filter":
			cli.AddNewFilterUI(config)
		case "del-filter":
			cli.DeleteFilterUI(config)
		case "clear-filters":
			cli.DeleteAllFiltersUI(config)
		case "del-elements":
			cli.DeleteElementsInFilterUI(config)
		case "add-elements":
			cli.AddElementsInFilterUI(config)
		case "info":
			cli.InformationAboutConfig(config)
		case "set-input-file":
			cli.ChangeInputFileUI(config)
		case "set-output-file":
			cli.ChangeOutputFileUI(config)
		case "help":
			cli.CommandsInformation()
		case "run":
			return
		default:
			fmt.Println("Введена неизвестная команда")
		}
	}
}
