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
	input, err := filepath.Abs(filepath.Clean(filepath.Join("data", "input")))
	if err != nil {
		return "", err
	}
	fmt.Printf("Убедитесь, что файл с исходными данными находится в папке input: %s\n", input)
	fmt.Println("")
	for {
		input, err := getReader(
			"Введите имя файла с исходными данными (например: data.csv): ",
		)
		if err != nil {
			return "", err
		}

		if filepath.Ext(input) != ".csv" {
			fmt.Println("Файл должен быть с расширением '.csv'.")
			continue
		}

		if strings.ContainsAny(input, `/\`) {
			fmt.Println("Введите только имя файла, без пути.")
			continue
		}
		inputPath, err := absClean(
			filepath.Join(r.BaseDir, "data", "input", input),
		)
		if err != nil {
			return "", err
		}

		if _, err := os.Stat(inputPath); err != nil {
			if os.IsNotExist(err) {
				fmt.Printf("Файла %s не существует, попробуйте другое имя.\n", input)
				continue
			}
			return "", err
		}

		return inputPath, nil
	}
}

// получение пути для выходного файла (после обработки)
func (r *Resolver) Output() (string, error) {
	for {
		output, err := getReader(
			"Введите имя файла в котором будут сохраняться данные после обработки (например: result.csv): ",
		)
		if err != nil {
			return "", err
		}
		if filepath.Ext(output) != ".csv" {
			fmt.Println("Файл должен быть с расширением '.csv'.")
			continue
		}

		if strings.ContainsAny(output, `/\`) {
			fmt.Println("Введите только имя файла, без пути.")
			continue
		}

		outputPath, err := absClean(
			filepath.Join(r.BaseDir, "data", "output", output),
		)

		if _, err = os.Stat(outputPath); err == nil {
			fmt.Printf("Файл %s уже существует, попробуйте другое имя.\n", filepath.Base(outputPath))
			continue
		}
		file, err := os.Create(outputPath)
		if err != nil {
			return "", err
		}
		defer file.Close()

		return outputPath, nil
	}
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
