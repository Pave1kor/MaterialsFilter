package config

import (
	pathFile "MaterialsFilter/internal/infrastructure/path"
	"fmt"
)

func ChangeInputFile(config *Config) error {
	path := pathFile.New("..")
	inputPath, err := path.Input()
	if err != nil {
		return err
	}
	config.Input = inputPath
	fmt.Println("Обрабатываемый файл изменен!")
	return nil
}
