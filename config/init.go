package config

import (
	"github.com/gdsc-ys/fluentify-server/src/handler"
	"github.com/gdsc-ys/fluentify-server/src/service"
)

type Initialization struct {
	userService service.UserService
	UserHandler handler.UserHandler
}

func NewInitialization(
	userService service.UserService,
	userHandler handler.UserHandler,
) *Initialization {
	return &Initialization{
		userService: userService,
		UserHandler: userHandler,
	}
}
