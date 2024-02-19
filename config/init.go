package config

import (
	"github.com/gdsc-ys/fluentify-server/src/handler"
	"github.com/gdsc-ys/fluentify-server/src/middleware"
	"github.com/gdsc-ys/fluentify-server/src/service"
)

type Initialization struct {
	AuthMiddleware middleware.AuthMiddleware

	UserService     service.UserService
	StorageService  service.StorageService
	TopicService    service.TopicService
	SentenceService service.SentenceService
	SceneService    service.SceneService

	UserHandler     handler.UserHandler
	TopicHandler    handler.TopicHandler
	SentenceHandler handler.SentenceHandler
	SceneHandler    handler.SceneHandler
	FeedbackHandler handler.FeedbackHandler
}

func NewInitialization(
	authMiddleware middleware.AuthMiddleware,

	userService service.UserService,
	storageService service.StorageService,
	topicService service.TopicService,
	sentenceService service.SentenceService,
	sceneService service.SceneService,

	userHandler handler.UserHandler,
	topicHandler handler.TopicHandler,
	sentenceHandler handler.SentenceHandler,
	sceneHandler handler.SceneHandler,
	feedbackHandler handler.FeedbackHandler,
) *Initialization {
	return &Initialization{
		AuthMiddleware:  authMiddleware,
		UserService:     userService,
		StorageService:  storageService,
		TopicService:    topicService,
		SentenceService: sentenceService,
		SceneService:    sceneService,
		UserHandler:     userHandler,
		TopicHandler:    topicHandler,
		SentenceHandler: sentenceHandler,
		SceneHandler:    sceneHandler,
		FeedbackHandler: feedbackHandler,
	}
}
