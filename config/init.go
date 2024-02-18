package config

import (
	"github.com/gdsc-ys/fluentify-server/src/handler"
	"github.com/gdsc-ys/fluentify-server/src/middleware"
	"github.com/gdsc-ys/fluentify-server/src/service"
)

type Initialization struct {
	AuthMiddleware middleware.AuthMiddleware

	UserService    service.UserService
	StorageService service.StorageService
	TopicService   service.TopicService

	UserHandler  handler.UserHandler
	TopicHandler handler.TopicHandler
}

func NewInitialization(
	authMiddleware middleware.AuthMiddleware,
	userService service.UserService,
	storageService service.StorageService,
	topicService service.TopicService,
	userHandler handler.UserHandler,
	topicHandler handler.TopicHandler,
) *Initialization {
	return &Initialization{
		AuthMiddleware: authMiddleware,
		UserService:    userService,
		StorageService: storageService,
		TopicService:   topicService,
		UserHandler:    userHandler,
		TopicHandler:   topicHandler,
	}
}
