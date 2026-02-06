package config

import (
	"fmt"
	"strings"
)

// Вывод в терминал информации о текущих настройках фильтров и входного файла
func OutputData(config *Config) {
	fmt.Println("Путь к входному файлу:", config.Input)
	if config.Filters == nil {
		fmt.Println("Фильтры отсутвуют! Добавьте их с помощью команды 'addF'.")
	}
	for _, filter := range config.Filters {
		fmt.Printf("Имя фильтра: %s\n", filter.Name)
		fmt.Printf("Путь к выходному файлу: %s\n", filter.Output)
		listElements(config, filter.Name)
	}
	fmt.Println()
}

// Вывод в терминал списка элементов выбранного фильтра
func listElements(config *Config, nameFilter string) {
	var elements []string
	for _, filter := range config.Filters {
		if filter.Name == nameFilter {
			for element := range filter.Filter {
				elements = append(elements, element)
			}
		}
	}
	fmt.Printf("Список элементов фильтра %s следующий: %s\n", nameFilter, strings.Join(elements, ","))
}

// Вывод в терминал доступных команд для изменения настроек фильтрации
func CommandsInformation() {
	fmt.Println("\nДоступные команды для изменения настроек фильтрации:")
	fmt.Println("addF    - добавить новый фильтр")
	fmt.Println("delF    - удалить существующий фильтр")
	fmt.Println("delAllF - удалить все фильтры")
	fmt.Println("addEl   - добавить элементы в существующий фильтр")
	fmt.Println("delEl   - удалить элементы из существующего фильтра")
	fmt.Println("run     - сохранить изменения в файл настроек и выйти из режима изменения")
	fmt.Println("info    - показать информацию о текущих настройках")
	fmt.Println("command - показать команды для изменения настроек фильтрации")
}

// В терминале выводится информация обо всех фильтрах, загруженных из конфигурационного файла
func informationAboutConfig(config *Config) {
	fmt.Println("\nФайл настроек.")
	fmt.Printf("Данные для обработки должны находиться по пути: %s\n", config.Input)

	if config.Filters == nil {
		fmt.Println("Фильтры отсутвуют! Добавьте их с помощью команды 'addF'.")
	}

	for _, filter := range config.Filters {
		fmt.Println()
		fmt.Printf("Имя Фильтра: %s\n", filter.Name)
		fmt.Printf("Обработанный файл будет находиться по пути: %s\n", filter.Output)
		var elements []string
		for element := range filter.Filter {
			elements = append(elements, element)
		}
		fmt.Printf("Элементы фильтра: %s\n", strings.Join(elements, ","))
	}
}
