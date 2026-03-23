package path

import (
	errorsx "MaterialsFilter/pkg/errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Задать путь к файлу конфигураций
func (path Base) SetConfigPathFile() (string, bool) {
	configPathFile := filepath.Join(path.configPathFolder, "config.json")
	_, err := os.Stat(configPathFile)
	return configPathFile, err == nil
}

// Задать путь к файлу с исходными данными
func SetInputPathFile(inputFileName, inputPathFolder string) (string, error) {
	inputPathFile := filepath.Join(inputPathFolder, inputFileName)

	if _, err := os.Stat(inputPathFile); err != nil {
		return "", fmt.Errorf("%w Измените имя файла.", errorsx.ErrInputNotExists)
	}
	return inputPathFile, nil
}

// Задать путь к файлу с результатами обработки
func SetOutputPathFile(outputFileName, outputPathFolder string) (string, error) {
	outputPathFile := filepath.Join(outputPathFolder, outputFileName)
	if _, err := os.Stat(outputPathFile); err == nil {
		return "", fmt.Errorf("%w Файл будет перезаписан.", errorsx.ErrFileExists)
	}
	file, err := os.Create(outputPathFile)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	return outputPathFile, nil
}

// Получение пути к папке с исходными данными
func (path Base) GetInputPathFolder() string {
	return path.inputPathFolder
}

// Получение пути к папке с результатами обработки
func (path Base) GetOutputPathFolder() string {
	return path.outputPathFolder
}
