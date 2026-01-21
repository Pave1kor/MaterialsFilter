package materials

type MaterialsName string
type MaterialsInformation map[MaterialsName]Material
type Material struct {
	InformationElements []string
	ParseMaterials      []string
	FilterHeusler       bool
	FilterChalcogenide  bool
}

func NewMaterials() MaterialsInformation {
	return make(MaterialsInformation)
}
