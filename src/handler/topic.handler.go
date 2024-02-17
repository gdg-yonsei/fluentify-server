package handler

import (
	"github.com/gdsc-ys/fluentify-server/src/model"
	"github.com/gdsc-ys/fluentify-server/src/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TopicHandler interface {
	ListTopics(c echo.Context) error
	GetTopic(c echo.Context) error
}

type TopicHandlerImpl struct {
	topicService   service.TopicService
	storageService service.StorageService
}

func (handler *TopicHandlerImpl) ListTopics(c echo.Context) error {
	topicSlice, err := handler.topicService.ListTopics()
	if err != nil {
		return model.NewCustomHTTPError(http.StatusInternalServerError, err)
	}

	if len(topicSlice) == 0 {
		return c.JSON(http.StatusOK, topicSlice)
	}

	for i, topic := range topicSlice {
		thumbnailURL, err := handler.storageService.GetFileUrl(topic.ThumbnailUrl)
		if err != nil {
			return model.NewCustomHTTPError(http.StatusInternalServerError, err)
		}
		topicSlice[i].ThumbnailUrl = thumbnailURL
	}

	return c.JSON(http.StatusOK, topicSlice)
}

func (handler *TopicHandlerImpl) GetTopic(c echo.Context) error {
	// TODO: Fix with protobuf
	id := c.FormValue("id")
	topic, err := handler.topicService.GetTopic(id)
	if err != nil {
		return model.NewCustomHTTPError(http.StatusInternalServerError, err)
	}

	thumbnailURL, err := handler.storageService.GetFileUrl(topic.ThumbnailUrl)
	if err != nil {
		return model.NewCustomHTTPError(http.StatusInternalServerError, err)
	}
	topic.ThumbnailUrl = thumbnailURL

	return c.JSON(http.StatusOK, topic)
}

func TopicHandlerInit(topicService service.TopicService, storageService service.StorageService) *TopicHandlerImpl {
	return &TopicHandlerImpl{
		topicService:   topicService,
		storageService: storageService,
	}
}
