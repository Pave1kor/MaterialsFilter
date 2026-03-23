package cli

import (
	cfg "MaterialsFilter/internal/domain/config"
	ptable "MaterialsFilter/internal/infrastructure/periodictable"
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/mattn/go-runewidth"
)

// Изменение списка элементов в фильтре
func changeListElements(existsFunc func(element string) error, changeFunc func(element string), checkFunc func() bool) {
	var newElement string

	fmt.Println("Изменение фильтра.")
	fmt.Println("Чтобы завершить ввод, оставьте строку пустой.")
	for {
		fmt.Print("Химический элемент: ")
		newElement = NewLine()

		if newElement == "" {
			if checkFunc() {
				break
			}
			fmt.Println("Фильтр пуст. Добавьте хотя бы один химический элемент.")
			continue
		}

		if !ptable.Get(newElement) {
			fmt.Println("Неизвестный химический элемент. Попробуйте снова.")
			continue
		}

		if err := existsFunc(newElement); err != nil {
			fmt.Println(err)
			continue
		}
		changeFunc(newElement)
	}
}

// Вывод в терминал списка элементов всех доступных фильтров
func listElementsInFilter(filters []cfg.Filter) {
	if len(filters) == 0 {
		fmt.Println("Список фильтров пуст.")
		return
	}
	fmt.Println("Список фильтров.")
	for val, filter := range filters {
		sort.Strings(filter.Elements)
		fmt.Printf("\nФильтр %d.\n", val+1)
		fmt.Printf("Имя фильтра: %s\n", filter.Name)
		fmt.Printf("Список элементов: %s\n", strings.Join(filter.Elements, ", "))
		fmt.Printf("Имя файла для сохранения результатов фильтрации: %s\n", filepath.Base(filter.OutputPathFile))
	}
	fmt.Println()
}

func toInterfaceSlice(s []string) []interface{} {
	r := make([]any, len(s))
	for i := range s {
		r[i] = s[i]
	}
	return r
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

// Ввод новой строки
func NewLine() string {
	reader := bufio.NewReader(os.Stdin)
	newLine, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	newLine = strings.TrimSpace(newLine)
	return newLine
}
