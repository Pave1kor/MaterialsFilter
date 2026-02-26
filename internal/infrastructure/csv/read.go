package csv

import (
	"encoding/csv"
	"io"
	"os"
	"strings"
)

// Чтение данных из csv файла
func (obj *CSVFile) ReadCSV() (map[string][]string, error) {
	m := make(map[string][]string)

	file, err := os.Open(obj.Input)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = obj.Comma

	record, err := r.Read()
	if err != nil {
		return nil, err
	}
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
		if len(record) > 6 {
			continue
		}
		m[record[0]] = record
	}

	for _, val := range m {
		obj.Table = val
		break
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
