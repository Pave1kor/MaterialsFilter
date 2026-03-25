package path

import (
	errorsx "MaterialsFilter/pkg/errors"
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
		return "", errorsx.ErrInputNotExists
	}
	return inputPathFile, nil
}

// Задать путь к файлу с результатами обработки
func SetOutputPathFile(outputFileName, outputPathFolder string) string {
	outputPathFile := filepath.Join(outputPathFolder, outputFileName)
	return outputPathFile
}

// Получение пути к папке с исходными данными
func (path Base) GetInputPathFolder() string {
	return path.inputPathFolder
}

// Получение пути к папке с результатами обработки
func (path Base) GetOutputPathFolder() string {
	return path.outputPathFolder
}

func (path Base) SyncOutputFileName() []string {
	var listFiles []string
	files, err := os.ReadDir(path.outputPathFolder)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		listFiles = append(listFiles, file.Name())
	}
	return listFiles
}
