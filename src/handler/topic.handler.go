package handler

import (
	pb "github.com/gdsc-ys/fluentify-server/gen/proto"
	"github.com/gdsc-ys/fluentify-server/src/converter"
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
	var request = pb.ListTopicsRequest{}
	if err := c.Bind(&request); err != nil {
		return model.NewCustomHTTPError(http.StatusBadRequest, err)
	}

	topicSlice, err := handler.topicService.ListTopics()
	if err != nil {
		return model.NewCustomHTTPError(http.StatusInternalServerError, err)
	}

	if len(topicSlice) == 0 {
		return c.JSON(http.StatusOK, topicSlice)
	}

	var compactTopicDTOs = make([]*pb.CompactTopicDTO, len(topicSlice))
	for i, topic := range topicSlice {
		thumbnailURL, err := handler.storageService.GetFileUrl(topic.ThumbnailUrl)
		if err != nil {
			return model.NewCustomHTTPError(http.StatusInternalServerError, err)
		}
		topic.ThumbnailUrl = thumbnailURL

		compactTopicDTOs[i] = converter.ToCompactTopicDTO(topic)
	}

	return c.JSON(http.StatusOK, compactTopicDTOs)
}

func (handler *TopicHandlerImpl) GetTopic(c echo.Context) error {
	var request = pb.GetTopicRequest{}
	if err := c.Bind(&request); err != nil {
		return model.NewCustomHTTPError(http.StatusBadRequest, err)
	}
	if request.GetId() == "" {
		return model.NewCustomHTTPError(http.StatusBadRequest, "topic id is required")
	}

	topic, err := handler.topicService.GetTopic(request.GetId())
	if err != nil {
		return model.NewCustomHTTPError(http.StatusInternalServerError, err)
	}

	thumbnailURL, err := handler.storageService.GetFileUrl(topic.ThumbnailUrl)
	if err != nil {
		return model.NewCustomHTTPError(http.StatusInternalServerError, err)
	}
	topic.ThumbnailUrl = thumbnailURL

	return c.JSON(http.StatusOK, converter.ToTopicDTO(topic))
}

func TopicHandlerInit(topicService service.TopicService, storageService service.StorageService) *TopicHandlerImpl {
	return &TopicHandlerImpl{
		topicService:   topicService,
		storageService: storageService,
	}
}
