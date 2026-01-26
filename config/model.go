package config

type Config struct {
	Filters []Filter
}

type Filter struct {
	Name   string
	Output string
	Input  string
	Filter string
}
