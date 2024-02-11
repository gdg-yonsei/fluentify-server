//go:build wireinject
// +build wireinject

package config

import (
	"github.com/gdsc-ys/fluentify-server/src/handler"
	"github.com/gdsc-ys/fluentify-server/src/middleware"
	"github.com/gdsc-ys/fluentify-server/src/service"
	"github.com/google/wire"
)

var firebaseApp = wire.NewSet(InitializeFirebaseApp)
var firebaseAuthClient = wire.NewSet(NewFirebaseAuthClient)
var authMiddlewareSet = wire.NewSet(wire.Struct(new(middleware.AuthMiddleware), "*"))
var userServiceSet = wire.NewSet(service.UserServiceInit, wire.Bind(new(service.UserService), new(*service.UserServiceImpl)))
var userHandlerSet = wire.NewSet(handler.UserHandlerInit, wire.Bind(new(handler.UserHandler), new(*handler.UserHandlerImpl)))

func Init() *Initialization {
	wire.Build(NewInitialization, firebaseApp, firebaseAuthClient, authMiddlewareSet, userServiceSet, userHandlerSet)
	return nil
}
