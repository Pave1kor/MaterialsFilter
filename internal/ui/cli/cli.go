package cli

import (
	cfg "MaterialsFilter/internal/domain/config"
	"fmt"
	"path/filepath"
	"strings"
)

// Добавить элемент в фильтр
func AddElementsInFilterUI(config *cfg.Config) error {
	fmt.Println()
	fmt.Println("Добавление новых элементов в фильтр.")
	listElementsInFilter(config.Filters)
	filterName, err := config.AddElementsInFilter()
	if err != nil {
		return err
	}
	fmt.Printf("Введенные элементы успешно добавлены в фильтр %s.\n", filterName)
	fmt.Println()
	return nil
}

// Добавить новый фильтр
func AddNewFilterUI(config *cfg.Config) error {
	fmt.Println()
	fmt.Println("Добавление нового фильтра.")
	listElementsInFilter(config.Filters)
	nameFilter, err := config.AddNewFilter()
	if err != nil {
		return err
	}
	fmt.Printf("Фильтр %s успешно создан!\n", nameFilter)
	return nil
}

// Удалить фильтр с заданным именем
func DeleteFilterUI(config *cfg.Config) error {
	fmt.Println()
	fmt.Println("Удаление фильтра по его имени.")
	listElementsInFilter(config.Filters)
	nameFilter, err := config.DeleteFilter()
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
	fmt.Println()
	fmt.Println("Удаление элементов из фильтра.")
	listElementsInFilter(config.Filters)
	err := config.DeleteElementsInFilter()
	if err != nil {
		return err
	}
	fmt.Println("Элементы успешно удалены из фильтра.")
	fmt.Println()
	return nil
}

// Вывод информации обо всех фильтрах, загруженных из ффайла настроек
func InformationAboutConfig(config cfg.Config) error {
	input, err := filepath.Abs(filepath.Clean(filepath.Join("data", "input")))
	if err != nil {
		return err
	}
	output, err := filepath.Abs(filepath.Clean(filepath.Join("data", "output")))
	if err != nil {
		return err
	}
	configFile, err := filepath.Abs(filepath.Clean(filepath.Join("configs", "config.json")))
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
	err := config.ChangeInputFile()
	if err != nil {
		return err
	}
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
	fmt.Println("changeIn - изменить имя файла с иходными данными")
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
		fmt.Printf("Список элементов фильтра %s следующий: %s\n", filter.Name, strings.Join(elements, ", "))
	}
}
