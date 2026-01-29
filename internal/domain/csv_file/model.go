package csvfile

type CSVFile struct {
	Input  string
	Output string
	Comma  rune
}

func NewCSVFile(input string, output string) *CSVFile {
	return &CSVFile{
		Input:  input,
		Output: output,
		Comma:  ';',
	}
}
