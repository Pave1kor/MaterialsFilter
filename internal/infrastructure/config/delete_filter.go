package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func deleteFilter(config *Config) error {

	reader := bufio.NewReader(os.Stdin)
	for {
		if config.Filters == nil {
			fmt.Println("Фильтры отсутвуют в файле настроек. Пожалуйста добавьте новый фильтр!")
			return nil
		}

		fmt.Print("Введите имя фильтра, который необходимо удалить: ")
		name, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		name = strings.TrimSpace(name)

		for i, filter := range config.Filters {
			if filter.Name == name {
				config.Filters = append(config.Filters[:i], config.Filters[i+1:]...)
				fmt.Printf("Фильтр %s успешно удален!", name)
				return nil
			}
		}
		fmt.Println("Введенного имени фильтра не найдено, попробуйте еще раз!")
	}
}
