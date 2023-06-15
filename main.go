package main

import (
	"github.com/testsabirweb/go-template/app"
	"github.com/testsabirweb/go-template/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
