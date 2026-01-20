package filter

// фильтр для проверки на сплав Гейслера
func HeuslerFilter() FilterLists {
	return FilterLists{
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
	}
}

// фильтр для проверки на халькогенид
func ChalcogenideFilter() FilterLists {
	return FilterLists{
		"W":  nil,
		"Bi": nil,
		"Sb": nil,
		"Te": nil,
		"Se": nil,
		"Mg": nil,
		"S":  nil,
	}
}
