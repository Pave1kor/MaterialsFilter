package app

import (
	cfg "MaterialsFilter/internal/domain/config"
	manager "MaterialsFilter/internal/domain/config_manager"
	json "MaterialsFilter/internal/infrastructure/json"
	pathFile "MaterialsFilter/internal/infrastructure/path"
	cli "MaterialsFilter/internal/ui/cli"
)

func Setup() (*cfg.Config, error) {

	// Инициализация каталогов
	err := pathFile.Path()
	if err != nil {
		return nil, err
	}

	// Инициализация пути конфигурационного файла
	configPath, err := pathFile.Config()
	if err != nil {
		return nil, err
	}

	// Получение пути расположения файла для обработки
	input, err := pathFile.Input()
	if err != nil {
		return nil, err
	}

	//  Создание пустого конфига
	config, err := LoadConfig(configPath, input)
	if err != nil {
		return nil, err
	}

	// Получение настроек из конфигурационного файла
	cli.InformationAboutConfig(*config)

	// Изменение настроек (добавление, удаление фильтров и т.д.)
	manager.ChangeConfig(config)

	// Сохранение настроек
	err = json.WriteJSON(configPath, *config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
