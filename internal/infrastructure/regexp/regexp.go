package filter

import "regexp"

func RegexpFilter(material string) ([]string, error) {
	re, err := regexp.Compile(`[A-Z][a-z]?`)
	if err != nil {
		return nil, err
	}
	return re.FindAllString(material, -1), nil
}
