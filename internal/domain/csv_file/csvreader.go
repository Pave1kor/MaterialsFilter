package csvfile

import (
	"encoding/csv"
	"io"
	"os"
)

// в дальнейшем можно вынести в infrastructure. Здесь создаем интерфейс
func (obj *CSVFile) ReadCSV() (map[string][]string, error) {
	file, err := os.Open(obj.Input)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = obj.Comma
	m := make(map[string][]string)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		m[record[0]] = record
	}
	return m, nil
}
