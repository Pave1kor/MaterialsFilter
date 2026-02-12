package main

import (
	app "MaterialsFilter/internal/app"
	"log"
)

func main() {
	config, err := app.Setup()
	if err != nil {
		log.Fatal(err)
	}
	app.Run(config)
}
