package main

import (
	app "MaterialsFilter/internal/app"
	cfg "MaterialsFilter/internal/domain/config"
	pathFile "MaterialsFilter/internal/infrastructure/path"
	"log"
)

func main() {
	// инициализация пути конфигурационного файла
	path := pathFile.New("..")
	configPath, err := path.Config()
	if err != nil {
		log.Fatal(err)
	}

	// инициализация путей
	err = path.Path()
	if err != nil {
		log.Fatal()
	}

	cfg.CheckConfig(configPath)              // создание конфигурационного файла, если его нет
	config, err := cfg.NewConfig(configPath) // получение конфигурации
	if err != nil {
		log.Fatal(err)
	}

	cfg.ConfigData(config)             // вывод информации о конфигурации
	cfg.ChangeConfig(config)           // изменение конфигурации (добавление, удаление фильтров и т.д.)
	cfg.SaveConfig(config, configPath) // сохранение конфигурации в файл
	app.Run(config)
}
