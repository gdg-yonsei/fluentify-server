package main

import (
	"github.com/gdsc-ys/fluentify-server/src/router"
)

func main() {
	echoR := router.Router()

	// Start server
	echoR.Logger.Fatal(echoR.Start(":8080"))
}
