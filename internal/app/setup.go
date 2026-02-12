package app

import (
	cfg "MaterialsFilter/internal/domain/config"
	manager "MaterialsFilter/internal/domain/config_manager"
	pathFile "MaterialsFilter/internal/infrastructure/path"
	"MaterialsFilter/internal/ui/cli"
)

func Setup() (*cfg.Config, error) {
	// инициализация пути конфигурационного файла
	path := pathFile.New("..")
	configPath, err := path.Config()
	if err != nil {
		return nil, err
	}

	// инициализация путей
	err = path.Path()
	if err != nil {
		return nil, err
	}

	config, err := cfg.NewConfig() //  Создание конфига с начальными значениями
	if err != nil {
		return nil, err
	}
	cli.InformationAboutConfig(*config)
	manager.ChangeConfig(config)        // изменение настроек (добавление, удаление фильтров и т.д.)
	config.SaveConfigToJSON(configPath) // сохранение конфигурации в файл
	return config, nil
}
