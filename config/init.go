package config

import (
	"github.com/gdsc-ys/fluentify-server/src/handler"
	"github.com/gdsc-ys/fluentify-server/src/middleware"
	"github.com/gdsc-ys/fluentify-server/src/service"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("unable to locate .env file")
	}
}

type Initialization struct {
	AuthMiddleware middleware.AuthMiddleware

	UserService    service.UserService
	StorageService service.StorageService

	UserHandler handler.UserHandler
}

func NewInitialization(
	authMiddleware middleware.AuthMiddleware,
	storageService service.StorageService,
	userService service.UserService,
	userHandler handler.UserHandler,
) *Initialization {
	return &Initialization{
		AuthMiddleware: authMiddleware,
		StorageService: storageService,
		UserService:    userService,
		UserHandler:    userHandler,
	}
}
