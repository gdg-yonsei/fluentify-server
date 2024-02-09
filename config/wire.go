//go:build wireinject
// +build wireinject

package config

import (
	"github.com/gdsc-ys/fluentify-server/src/handler"
	"github.com/gdsc-ys/fluentify-server/src/service"
	"github.com/google/wire"
)

var firebaseApp = wire.NewSet(InitializeFirebaseApp)

var userServiceSet = wire.NewSet(service.UserServiceInit, wire.Bind(new(service.UserService), new(*service.UserServiceImpl)))
var userHandlerSet = wire.NewSet(handler.UserHandlerInit, wire.Bind(new(handler.UserHandler), new(*handler.UserHandlerImpl)))

func Init() *Initialization {
	wire.Build(NewInitialization, firebaseApp, userServiceSet, userHandlerSet)
	return nil
}
