package config

// Добавление фильтра в конфиг
func (cfg *Config) AddFilterInConfig(newFilter Filter) {
	// newCfg := cfg.Filters
	// cfg.Filters = newCfg
	cfg.Filters = append(cfg.Filters, newFilter)
}

// Удаление фильтров
func (cfg *Config) DeleteAllFiltersFromConfig() {
	cfg.Filters = nil
}

// Извлечение фильтра из конфига
func (cfg *Config) ExtractFilterFromConfig(nameFilter string) (Filter, bool) {
	for idx := range cfg.Filters {
		if cfg.Filters[idx].Name == nameFilter {
			filter := cfg.Filters[idx]

			if idx < len(cfg.Filters)-1 {
				copy(cfg.Filters[idx:], cfg.Filters[idx+1:])
			}
			cfg.Filters = cfg.Filters[:len(cfg.Filters)-1]

			return filter, true
		}
	}
	return Filter{
		Name: nameFilter,
	}, false
}
