package config

import (
	"github.com/gdsc-ys/fluentify-server/src/handler"
	"github.com/gdsc-ys/fluentify-server/src/middleware"
	"github.com/gdsc-ys/fluentify-server/src/service"
)

type Initialization struct {
	AuthMiddleware middleware.AuthMiddleware
	UserService    service.UserService
	UserHandler    handler.UserHandler
}

func NewInitialization(
	authMiddleware middleware.AuthMiddleware,
	userService service.UserService,
	userHandler handler.UserHandler,
) *Initialization {
	return &Initialization{
		AuthMiddleware: authMiddleware,
		UserService:    userService,
		UserHandler:    userHandler,
	}
}
