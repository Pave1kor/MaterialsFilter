package config

// Добавление фильтра в конфиг
func (cfg *Config) AddFilterInConfig(filter *Filter) {
	cfg.Filters = append(cfg.Filters, *filter)
}

// Добавление имени файла с исходными данными в конфиг
func (cfg *Config) AddOutputFileNameInConfig(outputFileName string) {
	cfg.OutputFileNameList = append(cfg.OutputFileNameList, outputFileName)
}

// Удаление имени файла с исходными данными из конфига
func (cfg *Config) DeleteOutputFileNameFromConfig(outputFileName string) {
	newList := make([]string, 0, len(cfg.OutputFileNameList))

	for _, fileName := range cfg.OutputFileNameList {
		if fileName == outputFileName {
			continue
		}
		newList = append(newList, fileName)
	}
	cfg.OutputFileNameList = newList
}

// Удаление фильтров
func (cfg *Config) DeleteAllFiltersFromConfig() {
	cfg.Filters = nil
}

// Удаление фильтра
func (cfg *Config) DeleteFilterInConfig(nameFilter string) {
	newFilter := make([]Filter, 0, len(cfg.Filters))

	for _, filter := range cfg.Filters {
		if filter.Name == nameFilter {
			continue
		}
		newFilter = append(newFilter, filter)
	}
	cfg.Filters = newFilter
}

// Поиск фильтра в конфиге
func (cfg *Config) FindFilterInConfig(nameFilter string) (*Filter, bool) {
	for idx := range cfg.Filters {
		if cfg.Filters[idx].Name == nameFilter {
			return &cfg.Filters[idx], true
		}
	}
	return nil, false
}

// Обновление имени файла с исходными данными в конфиге
func (cfg *Config) UpdateOutputFilenameInConfig(oldOutputFileName, newOutputFileName string) {
	newList := make([]string, 0, len(cfg.OutputFileNameList))

	for _, fileName := range cfg.OutputFileNameList {
		if fileName == oldOutputFileName {
			newList = append(newList, newOutputFileName)
			continue
		}
		newList = append(newList, fileName)
	}
	cfg.OutputFileNameList = newList
}
