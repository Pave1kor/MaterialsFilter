package cli

import (
	cfg "MaterialsFilter/internal/domain/config"
	json "MaterialsFilter/internal/infrastructure/json"
	ptable "MaterialsFilter/internal/infrastructure/p_table"
	pathFile "MaterialsFilter/internal/infrastructure/path"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"
)

// Добавить элемент существующий фильтр
func AddElementsInFilterUI(config *cfg.Config) error {
	var nameFilter string
	var listElements []string
	var err error
	fmt.Println()
	fmt.Println("Добавление новых элементов в существующий фильтр.")
	listElementsInFilter(config.Filters)

	//Верификация фильтра
	nameFilter, err = verificationFilter(config, true)
	if err != nil {
		return err
	}

	// Добавление новых элементов
	listElements, err = createListElements()
	if err != nil {
		return err
	}
	config.AddElementsInFilter(nameFilter, listElements)
	fmt.Printf("Введенные элементы успешно добавлены в фильтр %s.\n", nameFilter)
	fmt.Println()
	return nil
}

func createListElements() ([]string, error) {
	var newElement string
	var listElements []string
	var err error
	fmt.Println("Создание списка химических элементов.")
	for {
		fmt.Println("Введите символ химического элемента (пустой ввод соответствует окончанию ввода):")
		newElement, err = NewLine()
		if err != nil {
			return nil, err
		}
		if newElement == "" {
			break
		}
		if !ptable.Get(newElement) {
			fmt.Println("Введен неизвестный химический элемент, попробуйте еще раз.")
			continue
		}

		if slices.Contains(listElements, newElement) {
			fmt.Println("Введенный химический элемент уже дообавлен в список, попробуйте еще раз.")
			continue
		}
		listElements = append(listElements, newElement)
	}
	return listElements, nil
}

func verificationFilter(config *cfg.Config, mustExist bool) (string, error) {
	var nameFilter string
	var err error
	fmt.Println("Выбор фильтра.")
	for {
		fmt.Print("Введите имя фильтра:")

		nameFilter, err = NewLine()
		if err != nil {
			return "", err
		}

		exist := config.ExistFilter(nameFilter)

		if mustExist && !exist {
			fmt.Println("Фильтра с введенным именем не существует, попробуйте еще раз.")
			continue
		}

		if !mustExist && exist {
			fmt.Println("Фильтр с введенным именем существует, попробуйте еще раз.")
			continue
		}

		confirmed, err := Verification()
		if err != nil {
			return "", err
		}

		if confirmed {
			return nameFilter, nil
		}
	}
}

// Ввод новой строки
func NewLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	newLine, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	newLine = strings.TrimSpace(newLine)
	return newLine, nil
}

// Проверка: да/нет
func Verification() (bool, error) {
	for {
		fmt.Println("Подтверждение ввода.")
		fmt.Print("Ожидаю (да/нет) или (yes/no): ")

		input, err := NewLine()
		if err != nil {
			return false, err
		}

		switch input {
		case "да", "yes", "y":
			return true, nil
		case "нет", "no", "n":
			return false, nil
		default:
			fmt.Println("Введите 'да' или 'нет'.")
		}
	}
}

// Добавить элементы в новый фильтр
func AddNewFilterUI(config *cfg.Config) error {
	var nameFilter string
	var listElements []string
	var err error
	fmt.Println()
	fmt.Println("Создание нового фильтра.")
	listElementsInFilter(config.Filters)

	nameFilter, err = verificationFilter(config, false)
	if err != nil {
		return err
	}

	output, err := pathFile.Output()
	if err != nil {
		return err
	}

	// Верификация введенных элементов
	listElements, err = createListElements()
	if err != nil {
		return err
	}

	config.AddNewFilter(output, nameFilter, listElements)
	fmt.Printf("Фильтр %s успешно создан!\n", nameFilter)
	return nil
}

// Удалить фильтр с заданным именем
func DeleteFilterUI(config *cfg.Config) error {
	var nameFilter string
	var err error
	fmt.Println()
	fmt.Println("Удаление фильтра по его имени.")
	listElementsInFilter(config.Filters)
	nameFilter, err = verificationFilter(config, true)
	if err != nil {
		return err
	}

	err = config.DeleteFilter(nameFilter)
	if err != nil {
		return err
	}

	fmt.Printf("Фильтр %s успешно удален!", nameFilter)
	fmt.Println()
	return nil

}

// Удалить все фильтры
func DeleteAllFiltersUI(config *cfg.Config) {
	fmt.Println()
	fmt.Println("Удаление всех фильтров.")
	config.DeleteAllFilters()
	fmt.Println("Все фильтры успешно удалены!")
	fmt.Println()
}

// Удалить элементы фильтра
func DeleteElementsInFilterUI(config *cfg.Config) error {
	var nameFilter string
	var deleteElements []string
	var err error

	fmt.Println()
	fmt.Println("Удаление элементов из фильтра.")
	listElementsInFilter(config.Filters)

	nameFilter, err = verificationFilter(config, true)
	if err != nil {
		return err
	}

	deleteElements, err = createListElements()
	if err != nil {
		return err
	}

	config.DeleteElementsInFilter(nameFilter, deleteElements)
	fmt.Println("Все элементы успешно удалены из фильтра.")
	fmt.Println()
	return nil
}

// Вывод информации обо всех фильтрах, загруженных из файла настроек
func InformationAboutConfig(config cfg.Config) error {
	exe, err := os.Executable()
	if err != nil {
		return err
	}
	baseDir := filepath.Dir(exe)
	input, err := filepath.Abs(filepath.Clean(filepath.Join(baseDir, "data", "input")))
	if err != nil {
		return err
	}
	output, err := filepath.Abs(filepath.Clean(filepath.Join(baseDir, "data", "output")))
	if err != nil {
		return err
	}
	configFile, err := filepath.Abs(filepath.Clean(filepath.Join(baseDir, "configs", "config.json")))
	if err != nil {
		return err
	}
	fmt.Println()
	fmt.Println("Информация о настройках.")
	fmt.Printf("Файл с исходными данными должен находиться по пути: %s\n", input)
	fmt.Printf("Результаты фильтрации будут располагаться по пути: %s\n", output)
	fmt.Printf("Файл настроек располагается по пути: %s\n", configFile)
	listElementsInFilter(config.Filters)
	fmt.Println()
	return nil
}

// Изменение имени обрабатываемого файла
func ChangeInputFileUI(config *cfg.Config) error {
	fmt.Println()
	fmt.Println("Изменение имени файла с исходными данными.")

	inputPath, err := pathFile.Input()
	if err != nil {
		return err
	}

	config.ChangeInputFile(inputPath)
	fmt.Println("Имя изменено!")
	return nil
}

// Вывод в терминал доступных команд для изменения настроек фильтрации
func CommandsInformation() {
	fmt.Println()
	fmt.Println("Доступные команды для изменения настроек фильтрации:")
	fmt.Println("addF     - добавить новый фильтр")
	fmt.Println("delF     - удалить существующий фильтр")
	fmt.Println("delAllF  - удалить все фильтры")
	fmt.Println("addEl    - добавить элементы в существующий фильтр")
	fmt.Println("delEl    - удалить элементы из существующего фильтра")
	fmt.Println("changeIn - изменить имя файла с исходными данными")
	fmt.Println("info     - показать текущие настройки")
	fmt.Println("command  - показать команды для изменения настроек фильтрации")
	fmt.Println("run      - сохранить изменения в файл настроек и выйти из режима изменения")
	fmt.Println()
}

// Вывод в терминал списка элементов всех доступных фильтров
func listElementsInFilter(filters []cfg.Filter) {
	for _, filter := range filters {
		var elements []string
		for element := range filter.Filter {
			elements = append(elements, element)
		}
		sort.Strings(elements)
		fmt.Printf("Список элементов фильтра %s следующий: %s\n", filter.Name, strings.Join(elements, ", "))
	}
}

func WriteJSONUI(config cfg.Config, configPath string, inputFunc func() (string, error)) error {
	fmt.Println()
	fmt.Println("Файл настроек отсутствует!")
	fmt.Println("Создание нового файла настроек.")
	input, err := inputFunc()
	if err != nil {
		return err
	}
	config.Input = input
	err = json.WriteJSON(configPath, config)
	if err != nil {
		return err
	}
	return nil
}
