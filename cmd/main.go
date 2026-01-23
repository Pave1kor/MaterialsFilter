package main

import (
	app "MaterialsFilter/internal/app"
	cfg "MaterialsFilter/internal/config"
)

func main() {
	cfg := cfg.Load()
	app.Run(cfg)
}
