package cli

import (
	cfg "MaterialsFilter/internal/domain/config"
	csvFile "MaterialsFilter/internal/infrastructure/csv"
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

	"github.com/mattn/go-runewidth"
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
	fmt.Println("Чтобы завершить ввод, оставьте строку пустой.")

	for {
		fmt.Print("Химический элемент: ")
		newElement, err = NewLine()
		if err != nil {
			return nil, err
		}
		if newElement == "" {
			break
		}
		if !ptable.Get(newElement) {
			fmt.Println("Неизвестный химический элемент. Попробуйте снова.")
			continue
		}

		if slices.Contains(listElements, newElement) {
			fmt.Println("Этот элемент уже добавлен. Попробуйте снова.")
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
		fmt.Print("Введите имя фильтра: ")

		nameFilter, err = NewLine()
		if err != nil {
			return "", err
		}
		nameFilter = strings.TrimSpace(nameFilter)
		if nameFilter == "" {
			fmt.Println("Имя фильтра не может быть пустым.")
			continue
		}

		exist := config.ExistFilter(nameFilter)

		if mustExist && !exist {
			fmt.Println("Фильтр с таким именем не существует. Попробуйте снова.")
			continue
		}

		if !mustExist && exist {
			fmt.Println("Фильтр с таким именем уже существует. Попробуйте снова.")
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
	fmt.Println()
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
	fmt.Println("Текущие настройки:")
	fmt.Println("------------------")
	fmt.Printf("Исходные данные: %s\n", input)
	fmt.Printf("Результаты фильтрации: %s\n", output)
	fmt.Printf("Файл настроек: %s\n", configFile)
	fmt.Printf("Загружено фильтров: %d\n", len(config.Filters))
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
	fmt.Println()
	return nil
}

// Вывод доступных команд режима изменения настроек
func CommandsInformation() {
	fmt.Println()
	fmt.Println("Доступные команды:")
	fmt.Println("------------------")
	fmt.Println("addF     — добавить новый фильтр")
	fmt.Println("delF     — удалить существующий фильтр")
	fmt.Println("delAllF  — удалить все фильтры")
	fmt.Println("addEl    — добавить элементы в фильтр")
	fmt.Println("delEl    — удалить элементы из фильтра")
	fmt.Println("changeIn — изменить имя файла с исходными данными")
	fmt.Println("info     — показать текущие настройки")
	fmt.Println("command  — показать эту справку")
	fmt.Println("run      — сохранить изменения и выйти")
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
	fmt.Println("Файл настроек создан.")
	fmt.Println()
	return nil
}

func ChangeHeadlinesUI(csv *csvFile.CSVFile) error {
	fmt.Println("\nИзменение заголовков столбцов.")
	viewTable(csv.Headlines, csv.Table)
	newHeadlines := make([]string, 0, len(csv.Headlines))
	fmt.Println("Введите новый заголовок (Enter — оставить как есть).")
	for _, headline := range csv.Headlines {
		fmt.Printf("%s -> ", headline)
		line, err := NewLine()
		if err != nil {
			return err
		}
		if line == "" {
			newHeadlines = append(newHeadlines, headline)
			continue
		}
		newHeadlines = append(newHeadlines, line)
	}
	csv.ChangeHeadlines(newHeadlines)
	fmt.Println("Все заголовки изменены.")
	return nil
}

// Преобразуем слайс строк в []interface{} для fmt.Printf
func toInterfaceSlice(s []string) []interface{} {
	r := make([]interface{}, len(s))
	for i := range s {
		r[i] = s[i]
	}
	return r
}

func viewTable(headlines []string, data []string) {

	widths := make([]int, len(headlines))
	all := append([]string{}, headlines...)
	if len(data) > 0 {
		all = append(all, data...)
	}

	for i := range headlines {
		maxWidth := 0
		for j := i; j < len(all); j += len(headlines) {
			w := runewidth.StringWidth(all[j])
			if w > maxWidth {
				maxWidth = w
			}
		}
		widths[i] = maxWidth
	}

	line := make([]string, len(widths))
	for i, w := range widths {
		line[i] = strings.Repeat("-", w)
	}
	border := "+-" + strings.Join(line, "-+-") + "-+"

	fmt.Println()
	fmt.Println(border)
	headLineStr := "|"
	for i, h := range headlines {
		headLineStr += " " + padRight(h, widths[i]) + " |"
	}
	fmt.Println(headLineStr)
	fmt.Println(border)

	// печать строки данных
	dataStr := "|"
	for i, d := range data {
		if isNumber(d) {
			dataStr += " " + padLeft(d, widths[i]) + " |"
		} else {
			dataStr += " " + padRight(d, widths[i]) + " |"
		}
	}
	fmt.Println(dataStr)
	fmt.Println(border)
}

func padRight(s string, width int) string {
	w := runewidth.StringWidth(s)
	if w >= width {
		return s
	}
	return s + strings.Repeat(" ", width-w)
}

func padLeft(s string, width int) string {
	w := runewidth.StringWidth(s)
	if w >= width {
		return s
	}
	return strings.Repeat(" ", width-w) + s
}

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return false
	}
	c := s[0]
	return (c >= '0' && c <= '9') || c == '-'
}
