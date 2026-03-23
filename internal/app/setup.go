package app

import (
	cfg "MaterialsFilter/internal/domain/config"
	manager "MaterialsFilter/internal/domain/configmanager"
	jsonFile "MaterialsFilter/internal/infrastructure/json"
	pathFile "MaterialsFilter/internal/infrastructure/path"
	cli "MaterialsFilter/internal/ui/cli"
)

func Setup() *cfg.Config {
	var (
		config cfg.Config
		exists bool
	)
	// Инициализация базовых каталогов
	path := pathFile.NewPath()

	// Задать путь к файлу конфигураций
	config.ConfigPathFile, exists = path.SetConfigPathFile()

	if exists {
		// Прочитать данные из конфига
		jsonFile.ReadJSON(&config)
	} else {
		// Получить путь к папке с исходными данными
		config.InputPathFolder = path.GetInputPathFolder()

		// Получить путь к папке с результатами обработки
		config.OutputPathFolder = path.GetOutputPathFolder()
		//Создать новый конфиг
		cli.WriteJSONUI(&config)
	}

	// Получение информации о настройках из конфигурационного файла
	cli.InformationAboutConfig(&config)

	// Изменение настроек (добавление, удаление фильтров и т.д.)
	manager.ChangeConfig(&config)

	// Сохранение настроек
	jsonFile.WriteJSON(&config)
	return &config
}
