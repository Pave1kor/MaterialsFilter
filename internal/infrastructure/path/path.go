package path

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// создание папок для хранения файлов
func Path() error {

	exe, err := os.Executable()
	if err != nil {
		return err
	}
	base := filepath.Dir(exe)
	testPath := filepath.Join(base, "configs")

	err = os.MkdirAll(testPath, 0755)
	if err != nil {
		base, _ = os.UserConfigDir()
	}

	dataPath := []string{
		filepath.Join(base, "configs"),
		filepath.Join(base, "data", "input"),
		filepath.Join(base, "data", "output"),
	}

	for _, dir := range dataPath {
		configPath, err := absClean(dir)
		if err != nil {
			return err
		}
		err = os.MkdirAll(configPath, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

// получение пути для хранения входного файла (исходный)
func Input() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	input, err := filepath.Abs(filepath.Clean(filepath.Join(filepath.Dir(exe), "data", "input")))
	if err != nil {
		return "", err
	}
	fmt.Printf("Поместите файл с исходными данными в папку input: %s\n", input)
	fmt.Println("")
	for {
		input, err := getReader(
			"Введите имя файла с исходными данными (например: data.csv): ",
		)
		if err != nil {
			return "", err
		}

		if strings.ToLower(filepath.Ext(input)) != ".csv" {
			fmt.Println("Файл должен быть с расширением '.csv'.")
			continue
		}

		if strings.ContainsAny(input, `/\`) {
			fmt.Println("Введите только имя файла, без пути.")
			continue
		}
		inputPath, err := absClean(
			filepath.Join(filepath.Dir(exe), "data", "input", input),
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
func Output() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	for {
		output, err := getReader(
			"Введите имя файла в котором будут сохраняться данные после обработки (например: result.csv): ",
		)
		if err != nil {
			return "", err
		}
		if strings.ToLower(filepath.Ext(output)) != ".csv" {
			fmt.Println("Файл должен быть с расширением '.csv'.")
			continue
		}

		if strings.ContainsAny(output, `/\`) {
			fmt.Println("Введите только имя файла, без пути.")
			continue
		}

		outputPath, err := absClean(
			filepath.Join(filepath.Dir(exe), "data", "output", output),
		)

		if _, err = os.Stat(outputPath); err == nil {
			fmt.Printf("Файл %s уже существует, попробуйте другое имя.\n", filepath.Base(outputPath))
			continue
		}
		file, err := os.Create(outputPath)
		if err != nil {
			return "", err
		}
		file.Close()

		return outputPath, nil
	}
}

// получение пути файла конфигурации
func Config() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	return absClean(filepath.Join(filepath.Dir(exe), "configs", "config.json"))
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
