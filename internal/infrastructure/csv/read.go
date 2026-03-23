package csv

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

// Чтение данных из csv файла
func (obj *CSVFile) ReadCSV() {
	data := [][]string{}

	file, err := os.Open(obj.Input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = obj.Comma

	record, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}
	record = trimEmptyTail(record)
	obj.Headlines = record

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		record = trimEmptyTail(record)
		if len(record) > 6 {
			continue
		}
		data = append(data, record)
	}
	obj.Data = data
	for _, val := range data {
		obj.Table = val
		break
	}
}

// Обрезаем хвостовые пустые элементы
func trimEmptyTail(s []string) []string {
	for len(s) > 0 && strings.TrimSpace(s[len(s)-1]) == "" {
		s = s[:len(s)-1]
	}
	return s
}
