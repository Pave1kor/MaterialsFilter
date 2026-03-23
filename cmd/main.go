package main

import (
	app "MaterialsFilter/internal/app"
)

func main() {
	config := app.Setup()
	app.Run(config)
}
