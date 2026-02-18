package config

type Config struct {
	Input   string `json:"input"`
	Filters []Filter
}

type Filter struct {
	Name   string              `json:"name"`
	Filter map[string]struct{} `json:"filter"`
	Output string              `json:"output"`
}
