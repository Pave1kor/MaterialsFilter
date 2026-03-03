package path

import (
	"bufio"
	errorsx "MaterialsFilter/pkg/errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Создание каталогов для хранения файлов input, output, config
func Path() error {

	base, err := os.Executable()
	if err != nil {
		return err
	}

	baseDir := filepath.Dir(base)
	dataPath := []string{
		filepath.Join(baseDir, "configs"),
		filepath.Join(baseDir, "data", "input"),
		filepath.Join(baseDir, "data", "output"),
	}

	for _, dir := range dataPath {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

// Получение пути файла с исходными данными
func Input(input string) (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}

	inputPath := filepath.Join(filepath.Dir(exe), "data", "input", input)
	if _, err := os.Stat(inputPath); err != nil {
		return "", err
	}

	return inputPath, nil
}

// Получение пути файла с результатами обработки
func Output(output string) (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}

	outputPath := filepath.Join(filepath.Dir(exe), "data", "output", output)

	if _, err = os.Stat(outputPath); err == nil {
		return "", errorsx.ErrFileExists
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}
	file.Close()

	return outputPath, nil
}

// Получение пути файла конфигурации
func Config() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	path := filepath.Join(filepath.Dir(exe), "configs", "config.json")
	return path, nil
}

func getReader(note string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(note)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = strings.TrimSpace(input)
	return input, nil
}
