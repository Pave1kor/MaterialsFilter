package app

import (
	cfg "MaterialsFilter/internal/domain/config"
	manager "MaterialsFilter/internal/domain/config_manager"
	json "MaterialsFilter/internal/infrastructure/json"
	pathFile "MaterialsFilter/internal/infrastructure/path"
	cli "MaterialsFilter/internal/ui/cli"
)

func Setup() (*cfg.Config, error) {
	// инициализация пути конфигурационного файла

	configPath, err := pathFile.Config()
	if err != nil {
		return nil, err
	}

	inputFunc := func() (string, error) {
		input, err := pathFile.Input()
		if err != nil {
			return "", err
		}
		return input, nil
	}

	// инициализация путей
	err = pathFile.Path()
	if err != nil {
		return nil, err
	}

	config, err := Load(configPath, inputFunc) //  Создание конфига с начальными значениями
	if err != nil {
		return nil, err
	}

	cli.InformationAboutConfig(*config)
	manager.ChangeConfig(config)              // изменение настроек (добавление, удаление фильтров и т.д.)
	err = json.WriteJSON(configPath, *config) // сохранение json
	if err != nil {
		return nil, err
	}
	return config, nil
}
