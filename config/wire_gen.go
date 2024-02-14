// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package config

import (
	"github.com/gdsc-ys/fluentify-server/src/handler"
	"github.com/gdsc-ys/fluentify-server/src/middleware"
	"github.com/gdsc-ys/fluentify-server/src/service"
	"github.com/google/wire"
)

// Injectors from wire.go:

func Init() *Initialization {
	app := InitializeFirebaseApp()
	client := NewFirebaseAuthClient(app)
	authMiddlewareImpl := middleware.AuthMiddlewareInit(client)
	storageClient := NewFirebaseStorageClient(app)
	storageServiceImpl := service.StorageServiceInit(storageClient)
	userServiceImpl := service.UserServiceInit(client)
	userHandlerImpl := handler.UserHandlerInit(userServiceImpl)
	initialization := NewInitialization(authMiddlewareImpl, storageServiceImpl, userServiceImpl, userHandlerImpl)
	return initialization
}

// wire.go:

var firebaseApp = wire.NewSet(InitializeFirebaseApp)

var firebaseAuthClient = wire.NewSet(NewFirebaseAuthClient)

var firebaseStorageClient = wire.NewSet(NewFirebaseStorageClient)

var authMiddlewareSet = wire.NewSet(middleware.AuthMiddlewareInit, wire.Bind(new(middleware.AuthMiddleware), new(*middleware.AuthMiddlewareImpl)))

var userServiceSet = wire.NewSet(service.UserServiceInit, wire.Bind(new(service.UserService), new(*service.UserServiceImpl)))

var storageServiceSet = wire.NewSet(service.StorageServiceInit, wire.Bind(new(service.StorageService), new(*service.StorageServiceImpl)))

var userHandlerSet = wire.NewSet(handler.UserHandlerInit, wire.Bind(new(handler.UserHandler), new(*handler.UserHandlerImpl)))