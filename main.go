package main

import (
	"github.com/gdsc-ys/fluentify-server/config"
	"github.com/gdsc-ys/fluentify-server/src/router"
)

func main() {
	init := config.Init()
	echoR := router.Router(init)

	// Start server
	echoR.Logger.Fatal(echoR.Start(":8080"))
}
