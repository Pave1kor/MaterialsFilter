package config

// type Unmarshaler interface {
// 	UnmarshalJSON([]byte) error
// }
// type Marshaler interface {
// 	MarshalJSON() ([]byte, error)
// }

type Config struct {
	Input   string `json:"input"`
	Filters []Filter
}

type Filter struct {
	Name   string            `json:"name"`
	Filter map[string]string `json:"filter"`
	Output string            `json:"output"`
}
