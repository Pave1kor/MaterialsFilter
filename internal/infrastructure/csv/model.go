package csv

type CSVFile struct {
	Input     string
	Headlines []string
	Table     []string
	Comma     rune
}

func NewCSVFile(input string) *CSVFile {
	return &CSVFile{
		Input: input,
		Comma: ';',
	}
}
