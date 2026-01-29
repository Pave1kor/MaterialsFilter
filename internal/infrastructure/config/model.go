package config

import (
	"encoding/json"
	"strings"
)

type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}

type Config struct {
	Filters []Filter `json: "filters"`
}

type Filter struct {
	Output string
	Input  string
	Filter map[string]struct{} `json: "-"`
}

type filterJSON struct {
	Output string `json:"output"`
	Input  string `json:"input"`
	Filter string `json:"filter"`
}

func (f *Filter) UnmarshalJSON(data []byte) error {
	var filterJSON filterJSON
	if err := json.Unmarshal(data, &filterJSON); err != nil {
		return err
	}

	f.Input = filterJSON.Input
	f.Output = filterJSON.Output

	if f.Filter == nil {
		f.Filter = make(map[string]struct{})
	}

	elementsArr := strings.Split(filterJSON.Filter, ",")
	for _, elements := range elementsArr {
		f.Filter[strings.Trim(elements, " ")] = struct{}{}
	}

	return nil
}
