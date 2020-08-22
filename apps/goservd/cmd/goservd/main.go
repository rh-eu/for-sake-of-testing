package main

import (
	"log"

	"github.com/rh/for-sake-of-testing/apps/goservd/pkg/app"
	"github.com/rh/for-sake-of-testing/apps/goservd/pkg/version"
)

func main() {
	app := app.NewApp()

	log.Printf("Starting goservd version: %v", version.VERSION)

	app.Run()
}
