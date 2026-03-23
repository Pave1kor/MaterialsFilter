package cli

import (
	pathFile "MaterialsFilter/internal/infrastructure/path"
	errorsx "MaterialsFilter/pkg/errors"
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

// Получение пути к файлу с исходными данными
func SetInputPathFileUI(inputPathFolder string) string {
	var (
		inputFileName string
		inputPathFile string
		err           error
	)
	fmt.Printf("Поместите файл с исходными данными в папку input.")
	fmt.Println("")
	fmt.Println("Введите имя файла с исходными данными (например: data.csv).")
	for {
		fmt.Print("Имя файла: ")
		inputFileName = NewLine()

		if strings.ToLower(filepath.Ext(inputFileName)) != ".csv" {
			fmt.Println("Файл должен быть с расширением '.csv'.")
			continue
		}

		err = verificationName(strings.TrimSuffix(inputFileName, ".csv"))
		if errors.Is(err, errorsx.ErrVerification) {
			continue
		}

		inputPathFile, err = pathFile.SetInputPathFile(inputFileName, inputPathFolder)
		if errors.Is(err, errorsx.ErrInputNotExists) {
			fmt.Printf("Файла %s не существует, попробуйте другое имя.\n", inputFileName)
			continue
		}

		if Verification() {
			return inputPathFile
		}
		fmt.Println("Окей, давай еще раз.")
	}
}

// Получение пути к файлу с результатами обработки
func SetOutputPathFileUI(outputPathFolder string) string {
	var (
		outputFileName string
		outputPathFile string
		err            error
	)
	fmt.Println("Пожалуйста, введите имя файла для сохранения результатов фильтрации (например: result.csv).")
	for {

		fmt.Print("Имя файла: ")
		outputFileName = NewLine()

		if strings.ToLower(filepath.Ext(outputFileName)) != ".csv" {
			fmt.Println("Файл должен быть с расширением '.csv'.")
			continue
		}

		err = verificationName(strings.TrimSuffix(outputFileName, ".csv"))
		if errors.Is(err, errorsx.ErrVerification) {
			continue
		}

		if !Verification() {
			fmt.Println("Окей, давай еще раз.")
			continue
		}

		outputPathFile, err = pathFile.SetOutputPathFile(outputFileName, outputPathFolder)
		if errors.Is(err, errorsx.ErrFileExists) {
			fmt.Printf("Файл %s уже существует, попробуйте другое имя.\n", outputFileName)
			continue
		}
		return outputPathFile
	}
}
