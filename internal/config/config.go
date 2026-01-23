package config

type Config struct {
	Paths   Paths
	Filters map[string]string
}

type Paths struct {
	InputData  string
	OutputData string
}

func Load() *Config {
	return &Config{
		Paths: Paths{
			InputData:  "../data/input/zt_value.csv",
			OutputData: "../data/output/",
		},
		Filters: map[string]string{
			"heusler":      "Fe, Ti, V, Cr, Mn, Co, Ni, Cu, Zr, Sn, Ga, Al, Si, Ge, W",
			"chalcogenide": "W, Bi, Sb, Te, Se, Mg, S",
		},
	}
}
