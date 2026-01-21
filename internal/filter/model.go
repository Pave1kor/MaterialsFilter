package filter

type Elements string
type FilterLists struct {
	values map[Elements]any
}

// фильтр для проверки на сплав Гейслера
func NewHeuslerFilter() *FilterLists {
	return &FilterLists{
		values: map[Elements]any{
			"Fe": nil,
			"Ti": nil,
			"V":  nil,
			"Cr": nil,
			"Mn": nil,
			"Co": nil,
			"Ni": nil,
			"Cu": nil,
			"Zr": nil,
			"Sn": nil,
			"Ga": nil,
			"Al": nil,
			"Si": nil,
			"Ge": nil,
			"W":  nil,
		},
	}
}

// фильтр для проверки на халькогенид
func NewChalcogenideFilter() *FilterLists {
	return &FilterLists{
		values: map[Elements]any{
			"W":  nil,
			"Bi": nil,
			"Sb": nil,
			"Te": nil,
			"Se": nil,
			"Mg": nil,
			"S":  nil,
		},
	}
}
