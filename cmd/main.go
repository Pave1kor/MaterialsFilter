package main

import (
	cfg "MaterialsFilter/config"
	app "MaterialsFilter/internal/app"
	"log"
)

func main() {
	cfg, err := cfg.Load()
	if err != nil {
		log.Fatal(err)
	}
	app.Run(cfg)
}
