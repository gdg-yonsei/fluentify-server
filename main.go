package main

import (
	"fluentify-server/src/router"
)

func main() {
	echoR := router.Router()

	// Start server
	echoR.Logger.Fatal(echoR.Start(":8080"))
}
