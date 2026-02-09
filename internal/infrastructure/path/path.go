package path

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// создание папок для хранения файлов
func (r *Resolver) Path() error {
	// config
	configPath, err := absClean(filepath.Join(r.BaseDir, "configs"))
	if err != nil {
		return err
	}
	err = os.MkdirAll(configPath, os.ModePerm)
	if err != nil {
		return err
	}
	//input
	inputPath, err := absClean(filepath.Join(r.BaseDir, "data", "input"))
	if err != nil {
		return err
	}
	err = os.MkdirAll(inputPath, os.ModePerm)
	if err != nil {
		return err
	}
	// output
	outputPath, err := absClean(filepath.Join(r.BaseDir, "data", "output"))
	if err != nil {
		return err
	}
	err = os.MkdirAll(outputPath, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// получение пути для хранения входного файла (исходный)
func (r *Resolver) Input() (string, error) {
	for {
		input, err := getReader("Введите имя обрабатываемого файла \n (файл должен располгаться по пути: `data/input/input.csv`):")
		if err != nil {
			return "", err
		}

		inputPath, err := absClean(filepath.Join(r.BaseDir, "data", "input", input))
		if err != nil {
			return "", err
		}
		_, err = os.Stat(inputPath)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Printf("Файла с именем %s не существует, попробуйте еще раз\n", filepath.Base(inputPath))
				continue
			}
		}
		return inputPath, nil
	}
}

// получение пути для выходного файла (после обработки)
func (r *Resolver) Output() (string, error) {
	output, err := getReader("Введите имя выходного файла  (файл должен располгаться по пути: `data/input/input.csv`):")
	if err != nil {
		return "", err
	}

	return absClean(filepath.Join(r.BaseDir, "data", "output", output))
}

// получение пути файла конфигурации
func (r *Resolver) Config() (string, error) {
	return absClean(filepath.Join(r.BaseDir, "configs", "config.json"))
}

func absClean(s string) (string, error) {
	return filepath.Abs(filepath.Clean(s))
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
