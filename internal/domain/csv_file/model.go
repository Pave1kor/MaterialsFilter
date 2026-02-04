package csvfile

type CSVFile struct {
	Input  string
	Output string
	Comma  rune
}

func NewCSVFile(input string) *CSVFile {
	return &CSVFile{
		Input: input,
		Comma: ';',
	}
}
