package main

import (
	app "MaterialsFilter/internal/app"
	cfg "MaterialsFilter/internal/infrastructure/config"
	"log"
)

func main() {
	cfg, err := cfg.Load()
	if err != nil {
		log.Fatal(err)
	}
	app.Run(cfg)
}
