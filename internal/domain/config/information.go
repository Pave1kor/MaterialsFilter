package config

import (
	"fmt"
	"strings"
)

// Вывод в терминал информации о текущих настройках фильтров и входного файла
func ConfigData(config *Config) {
	fmt.Println()
	fmt.Println("Путь к входному файлу:", config.Input)

	for _, filter := range config.Filters {
		fmt.Printf("Имя фильтра: %s\n", filter.Name)
		fmt.Printf("Путь к выходному файлу: %s\n", filter.Output)
		listElements(config.Filters, filter.Name)
	}
	fmt.Println()
}

// Вывод в терминал списка элементов выбранного фильтра
func listElements(filters []Filter, nameFilter string) {
	fmt.Println()
	var elements []string
	for _, filter := range filters {
		if filter.Name == nameFilter {
			for element := range filter.Filter {
				elements = append(elements, element)
			}
		}
	}
	fmt.Printf("Список элементов фильтра %s следующий: %s\n", nameFilter, strings.Join(elements, ","))
	fmt.Println()
}

// Вывод в терминал доступных команд для изменения настроек фильтрации
func commandsInformation() {
	fmt.Println()
	fmt.Println("\nДоступные команды для изменения настроек фильтрации:")
	fmt.Println("addF     - добавить новый фильтр")
	fmt.Println("delF     - удалить существующий фильтр")
	fmt.Println("delAllF  - удалить все фильтры")
	fmt.Println("addEl    - добавить элементы в существующий фильтр")
	fmt.Println("delEl    - удалить элементы из существующего фильтра")
	fmt.Println("run      - сохранить изменения в файл настроек и выйти из режима изменения")
	fmt.Println("changeIn - изменить имя файла с входными данными")
	fmt.Println("info     - показать информацию о текущих настройках")
	fmt.Println("command  - показать команды для изменения настроек фильтрации")
	fmt.Println()
}

// В терминале выводится информация обо всех фильтрах, загруженных из конфигурационного файла
func informationAboutConfig(config *Config) {
	fmt.Println()
	fmt.Println("\nФайл настроек.")
	fmt.Printf("Данные для обработки должны находиться по пути: %s\n", config.Input)

	for _, filter := range config.Filters {
		fmt.Printf("Имя Фильтра: %s\n", filter.Name)
		fmt.Printf("Обработанный файл будет находиться по пути: %s\n", filter.Output)
		var elements []string
		for element := range filter.Filter {
			elements = append(elements, element)
		}
		fmt.Printf("Элементы фильтра: %s\n", strings.Join(elements, ","))
	}
	fmt.Println()
}
