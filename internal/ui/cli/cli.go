package cli

import (
	cfg "MaterialsFilter/internal/domain/config"
	csvFile "MaterialsFilter/internal/infrastructure/csv"
	jsonFile "MaterialsFilter/internal/infrastructure/json"
	errorsx "MaterialsFilter/pkg/errors"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/mattn/go-runewidth"
)

// Добавить элемент в существующий фильтр
func AddElementsInFilterUI(config *cfg.Config) {
	var (
		nameFilter string
		filter     cfg.Filter
		existFunc  func(string) error
		addFunc    func(string)
		checkFunc  func() bool
		err        error
		found      bool
	)
	fmt.Println()
	fmt.Println("Добавление новых элементов в существующий фильтр.")
	listElementsInFilter(config.Filters)

	// Найти существующий фильтр
	for {
		fmt.Print("Имя фильтра из списка: ")
		nameFilter = NewLine()
		err = verificationName(nameFilter)
		if errors.Is(err, errorsx.ErrVerification) {
			continue
		}
		filter, found = config.ExtractFilterFromConfig(nameFilter)
		if found {
			break
		}
		fmt.Println("Фильтра с таким именем не существует. Попробуйте снова.")
	}

	// Проверка: существует ли химический элемент в фильтре
	existFunc = func(element string) error {
		if slices.Contains(filter.Elements, element) {
			return errorsx.ErrElementExists
		}
		return nil
	}

	// Добавление элемента в фильтр
	addFunc = func(element string) {
		filter.Elements = append(filter.Elements, element)
	}

	// Проверка: имеются ли записи в фильтре
	checkFunc = func() bool {
		return len(filter.Elements) > 0
	}

	// Добавление новых элементов в существующий фильтр
	changeListElements(existFunc, addFunc, checkFunc)

	// Добавление фильтра в конфиг
	config.AddFilterInConfig(filter)
	fmt.Printf("Введенные элементы успешно добавлены в фильтр %s.\n", nameFilter)
	fmt.Println()
}

// Добавить элементы в новый фильтр
func AddNewFilterUI(config *cfg.Config) {
	var (
		filter     cfg.Filter
		nameFilter string
		existFunc  func(string) error
		addFunc    func(string)
		checkFunc  func() bool
		err        error
		found      bool
	)
	fmt.Println()
	fmt.Println("Создание нового фильтра.")
	listElementsInFilter(config.Filters)

	// Создать фильтр с уникальным именем
	for {
		fmt.Print("Имя нового фильтра: ")
		nameFilter = NewLine()
		err = verificationName(nameFilter)
		if errors.Is(err, errorsx.ErrVerification) {
			continue
		}
		filter, found = config.ExtractFilterFromConfig(nameFilter)
		if !found {
			break
		}
		fmt.Println("Фильтр с таким именем уже существует. Попробуйте снова.")
	}

	// Ввод имени файла
	filter.OutputPathFile = SetOutputPathFileUI(config.OutputPathFolder)

	// Проверка: существует ли химический элемент в фильтре
	existFunc = func(element string) error {
		if slices.Contains(filter.Elements, element) {
			return errorsx.ErrElementExists
		}
		return nil
	}

	// Добавить элемент в фильтр
	addFunc = func(element string) {
		filter.Elements = append(filter.Elements, element)
	}

	// Проверка: имеются ли записи в фильтре
	checkFunc = func() bool {
		return len(filter.Elements) > 0
	}

	// Создание списка химических элементов элементов
	changeListElements(existFunc, addFunc, checkFunc)

	// Добавление нового фильтра в конфиг
	config.AddFilterInConfig(filter)
	fmt.Printf("Фильтр %s успешно создан!\n", filter.Name)
	fmt.Println()
}

// Удалить фильтр с заданным именем
func DeleteFilterUI(config *cfg.Config) {
	var (
		nameFilter string
		found      bool
		err        error
	)
	fmt.Println()
	fmt.Println("Удаление фильтра по его имени.")
	listElementsInFilter(config.Filters)

	// Извлечь фильтр из конфига
	for {
		fmt.Print("Имя фильтра, который вы желаете удалить: ")
		nameFilter = NewLine()

		err = verificationName(nameFilter)
		if errors.Is(err, errorsx.ErrVerification) {
			continue
		}
		_, found = config.ExtractFilterFromConfig(nameFilter)

		if found {
			break
		}
		fmt.Println("Фильтра с таким именем не существует. Попробуйте снова.")
	}

	fmt.Printf("Фильтр %s успешно удален!", nameFilter)
	fmt.Println()
}

// Удалить все фильтры
func DeleteAllFiltersUI(config *cfg.Config) {
	fmt.Println()
	fmt.Println("Удаление всех фильтров.")
	config.DeleteAllFiltersFromConfig()
	fmt.Println("Все фильтры успешно удалены!")
	fmt.Println()
}

// Удалить элементы фильтра
func DeleteElementsInFilterUI(config *cfg.Config) {
	var (
		err        error
		filter     cfg.Filter
		existFunc  func(string) error
		delFunc    func(string)
		checkFunc  func() bool
		nameFilter string
		found      bool
	)

	fmt.Println()
	fmt.Println("Удаление элементов из фильтра.")
	listElementsInFilter(config.Filters)

	// Поиск и извлечение существующего фильтра по его имени
	for {
		fmt.Print("Имя фильтра из списка: ")
		nameFilter = NewLine()
		err = verificationName(nameFilter)
		if errors.Is(err, errorsx.ErrVerification) {
			continue
		}
		filter, found = config.ExtractFilterFromConfig(nameFilter)
		if found {
			break
		}
		fmt.Println("Фильтра с таким именем не существует. Попробуйте снова.")
	}

	// Проверка: существует ли химический элемент в фильтре.
	existFunc = func(element string) error {
		if slices.Contains(filter.Elements, element) {
			return nil
		}
		return errorsx.ErrElementNotExists
	}

	// Удаление элемента из фильтра
	delFunc = func(element string) {
		if len(filter.Elements) > 1 {
			filter.Elements = slices.DeleteFunc(filter.Elements, func(i string) bool {
				return i == element
			})
		} else {
			fmt.Println("В фильтре остался последний химический элемент. Удаление элемента невозможно.")
		}
	}

	// Проверка: имеются ли записи в фильтре
	checkFunc = func() bool {
		return len(filter.Elements) > 0
	}

	// Создание списка химических элементов элементов
	changeListElements(existFunc, delFunc, checkFunc)

	// Добавление нового фильтра в конфиг
	config.AddFilterInConfig(filter)
	fmt.Println("Все элементы успешно удалены из фильтра.")
	fmt.Println()
}

// Вывод информации обо всех фильтрах, загруженных из файла настроек
func InformationAboutConfig(config *cfg.Config) {

	fmt.Println()
	fmt.Println("Текущие настройки:")
	fmt.Println("------------------")
	fmt.Printf("Файл с исходными данными: %s\n", config.InputPathFile)
	fmt.Printf("Файл настроек: %s\n", config.ConfigPathFile)
	fmt.Printf("Результаты фильтрации: %s\n", config.OutputPathFolder)
	if len(config.Filters) > 0 {
		fmt.Printf("Загружено фильтров: %d\n", len(config.Filters))
	}
	listElementsInFilter(config.Filters)
	fmt.Println()

}

// Вывод в терминал доступных команд режима изменения настроек
func CommandsInformation() {
	fmt.Println()
	fmt.Println("Доступные команды:")
	fmt.Println("------------------")
	fmt.Println("add-filter       — добавить новый фильтр")
	fmt.Println("del-filter       — удалить существующий фильтр")
	fmt.Println("clear-filters    — удалить все фильтры")
	fmt.Println("add-elements     — добавить элементы в фильтр")
	fmt.Println("del-elements     — удалить элементы из фильтра")
	fmt.Println("set-input-file   — изменить имя файла с исходными данными")
	fmt.Println("set-output-file  — изменить имя файла с данными после фильтрации")
	fmt.Println("info             — показать текущие настройки")
	fmt.Println("help             — показать эту справку")
	fmt.Println("run              — сохранить изменения и выйти")
	fmt.Println()
}

// Создание файла настроек (пользовательский интерфейс)
func WriteJSONUI(config *cfg.Config) {
	fmt.Println()
	fmt.Println("Файл настроек отсутствует!")
	fmt.Println("Создание нового файла настроек.")

	config.InputPathFile = SetInputPathFileUI(config.InputPathFolder)
	jsonFile.WriteJSON(config)

	fmt.Println("Файл настроек создан.")
	fmt.Println()
}

// Изменение заголовков столбцов
func ChangeHeadlinesUI(csv *csvFile.CSVFile) error {
	fmt.Println("\nИзменение заголовков столбцов.")
	newHeadlines := make([]string, 0, len(csv.Headlines))
	fmt.Println("Введите новый заголовок (Enter — оставить как есть).")
	for _, headline := range csv.Headlines {
		fmt.Printf("%s -> ", headline)
		line := NewLine()

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

// Вывод в терминал таблицы с данными (шапка таблицы и первая строка)
func ViewTable(headlines []string, data []string) {

	widths := make([]int, len(headlines))
	for i := range widths {
		line1 := runewidth.StringWidth(headlines[i])
		line2 := runewidth.StringWidth(data[i])
		widths[i] = max(line1, line2)
	}

	line := make([]string, len(widths))
	for i, w := range widths {
		line[i] = strings.Repeat("-", w)
	}
	border := "+-" + strings.Join(line, "-+-") + "-+"

	fmt.Println()
	fmt.Println(border)
	var headLineStr strings.Builder
	headLineStr.WriteString("|")
	for i, h := range headlines {
		headLineStr.WriteString(" " + padRight(h, widths[i]) + " |")
	}
	fmt.Println(headLineStr.String())
	fmt.Println(border)

	var dataStr strings.Builder
	dataStr.WriteString("|")
	for i, d := range data {
		if isNumber(d) {
			dataStr.WriteString(" " + padLeft(d, widths[i]) + " |")
		} else {
			dataStr.WriteString(" " + padRight(d, widths[i]) + " |")
		}
	}
	fmt.Println(dataStr.String())
	fmt.Println(border)
}

// Изменение имени файла для сохранения результатов обработки
func ChangeOutputFileUI(config *cfg.Config) error {
	var (
		nameFilter string
		err        error
		filter     cfg.Filter
		found      bool
	)
	fmt.Println("\nИзменение имени файла для сохранения результатов фильтрации.")
	listElementsInFilter(config.Filters)

	for {
		fmt.Print("Имя фильтра: ")
		nameFilter = NewLine()
		err = verificationName(nameFilter)
		if errors.Is(err, errorsx.ErrVerification) {
			continue
		}
		filter, found = config.ExtractFilterFromConfig(nameFilter)
		if found {
			break
		}
		fmt.Println("Фильтра с таким именем не существует. Попробуйте снова.")
	}
	filter.OutputPathFile = SetOutputPathFileUI(config.OutputPathFolder)

	config.AddFilterInConfig(filter)
	return nil
}

// Изменение имени обрабатываемого файла
func ChangeInputFileUI(config *cfg.Config) error {
	fmt.Println()
	fmt.Println("Изменение имени файла с исходными данными.")

	config.InputPathFile = SetInputPathFileUI(config.InputPathFolder)

	fmt.Println("Имя изменено!")
	fmt.Println()
	return nil
}
