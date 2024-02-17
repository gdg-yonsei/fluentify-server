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
var firebaseStorageClient = wire.NewSet(NewFirebaseStorageClient)
var fireStoreClient = wire.NewSet(NewFireStoreClient)
var authMiddlewareSet = wire.NewSet(middleware.AuthMiddlewareInit, wire.Bind(new(middleware.AuthMiddleware), new(*middleware.AuthMiddlewareImpl)))
var userServiceSet = wire.NewSet(service.UserServiceInit, wire.Bind(new(service.UserService), new(*service.UserServiceImpl)))
var storageServiceSet = wire.NewSet(service.StorageServiceInit, wire.Bind(new(service.StorageService), new(*service.StorageServiceImpl)))
var topicServiceSet = wire.NewSet(service.TopicServiceInit, wire.Bind(new(service.TopicService), new(*service.TopicServiceImpl)))
var userHandlerSet = wire.NewSet(handler.UserHandlerInit, wire.Bind(new(handler.UserHandler), new(*handler.UserHandlerImpl)))
var topicHandlerSet = wire.NewSet(handler.TopicHandlerInit, wire.Bind(new(handler.TopicHandler), new(*handler.TopicHandlerImpl)))

func Init() *Initialization {
	wire.Build(
		NewInitialization,
		firebaseApp,
		firebaseAuthClient,
		firebaseStorageClient,
		fireStoreClient,
		authMiddlewareSet,
		userServiceSet,
		storageServiceSet,
		topicServiceSet,
		userHandlerSet,
		topicHandlerSet,
	)
	return nil
}
