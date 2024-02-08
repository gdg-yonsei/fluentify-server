package main

import (
	"github.com/gdsc-ys/fluentify-server/src/router"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("unable to locate .env file")
	}
}

func main() {
	echoR := router.Router()
	// Start server
	echoR.Logger.Fatal(echoR.Start(":8080"))
}
