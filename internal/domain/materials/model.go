package materials

type MaterialsName string
type MaterialsInformation map[MaterialsName]MaterialRow
type MaterialRow struct {
	InformationElements []string
	ParseMaterials      []string
	FiltersName         map[string]bool
}

func NewMaterials() MaterialsInformation {
	return make(MaterialsInformation, 1000)
}
func (m *MaterialRow) NewMaterialsRow() {
	if m.FiltersName == nil {
		m.FiltersName = make(map[string]bool)
	}
}
