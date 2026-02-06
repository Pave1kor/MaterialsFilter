package path

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// получение пути для хранения входного файла (исходный)
func (r *Resolver) Input() (string, error) {
	input, err := getReader("Введите имя входного файла (например, input.csv):")
	if err != nil {
		return "", err
	}
	return absClean(filepath.Join(r.BaseDir, "data", "input", input))
}

// получение пути для выходного файла (после обработки)
func (r *Resolver) Output() (string, error) {
	output, err := getReader("Введите имя выходного файла (например, output.csv):")
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
