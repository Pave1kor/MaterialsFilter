package config

type Config struct {
	InputPathFolder    string   `json:"inputPathFolder"`
	InputPathFile      string   `json:"inputPathFile"`
	OutputPathFolder   string   `json:"outputPathFolder"`
	OutputFileNameList []string `json:"outputPathFile"`
	ConfigPathFile     string   `json:"configPathFile"`
	Filters            []Filter
}

type Filter struct {
	Name           string   `json:"name"`
	Elements       []string `json:"filter"`
	OutputPathFile string   `json:"outputPathFile"`
}
