package path

import (
	"log"
	"os"
	"path/filepath"
)

type Base struct {
	configPathFolder string
	inputPathFolder  string
	outputPathFolder string
}

func NewPath() *Base {
	base, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	baseDir := filepath.Dir(base)
	dataPath := []string{
		filepath.Join(baseDir, "configs"),
		filepath.Join(baseDir, "data", "input"),
		filepath.Join(baseDir, "data", "output"),
	}

	for _, dir := range dataPath {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
	return &Base{
		inputPathFolder:  dataPath[1],
		outputPathFolder: dataPath[2],
		configPathFolder: dataPath[0],
	}
}
