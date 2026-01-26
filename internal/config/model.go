package config

type Config struct {
	Paths   []Path
	Filters []Filter
}

type Path struct {
	InputData string
}
type Filter struct {
	NameFilter string
	OutputData string
	DataFilter string
}
