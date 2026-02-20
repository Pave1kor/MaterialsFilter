package csv

import (
	"encoding/csv"
	"io"
	"os"
	"strings"
)

// в дальнейшем можно вынести в infrastructure. Здесь создаем интерфейс
func (obj *CSVFile) ReadCSV() (map[string][]string, error) {
	headBool := true
	file, err := os.Open(obj.Input)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = obj.Comma
	m := make(map[string][]string)
	record, err := r.Read()
	record = trimEmptyTail(record)
	obj.Headlines = record
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		record = trimEmptyTail(record)

		if headBool {
			obj.Table = record
			headBool = false
		}

		m[record[0]] = record
	}
	return m, nil
}

// Обрезаем хвостовые пустые элементы
func trimEmptyTail(s []string) []string {
	for len(s) > 0 && strings.TrimSpace(s[len(s)-1]) == "" {
		s = s[:len(s)-1]
	}
	return s
}
