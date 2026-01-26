package main

import (
	app "MaterialsFilter/internal/app"
	cfg "MaterialsFilter/internal/config"
)

func main() {
	cfg, err := cfg.Load()
	if err != nil {
		return
	}
	app.Run(cfg)
}
