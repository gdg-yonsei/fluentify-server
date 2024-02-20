package handler

import (
	pb "github.com/gdsc-ys/fluentify-server/gen/proto"
	"github.com/gdsc-ys/fluentify-server/src/converter"
	"github.com/gdsc-ys/fluentify-server/src/model"
	"github.com/gdsc-ys/fluentify-server/src/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type SentenceHandler interface {
	GetSentence(c echo.Context) error
}

type SentenceHandlerImpl struct {
	sentenceService service.SentenceService
	storageService  service.StorageService
}

func (handler *SentenceHandlerImpl) GetSentence(c echo.Context) error {
	var request = pb.GetSentenceRequest{}
	if err := c.Bind(&request); err != nil {
		return model.NewCustomHTTPError(http.StatusBadRequest, err)
	}
	if request.GetId() == "" {
		return model.NewCustomHTTPError(http.StatusBadRequest, "sentence id is required")
	}

	sentence, err := handler.sentenceService.GetSentence(request.GetId())
	if err != nil {
		return model.NewCustomHTTPError(http.StatusInternalServerError, err)
	}

	exampleVideoUrl, err := handler.storageService.GetFileUrl(sentence.VideoPath)
	if err != nil {
		return model.NewCustomHTTPError(http.StatusInternalServerError, err)
	}

	response := &pb.GetSentenceResponse{Sentence: converter.ToSentenceDTO(sentence, exampleVideoUrl)}
	return c.JSON(http.StatusOK, response)
}

func SentenceHandlerInit(sentenceService service.SentenceService) *SentenceHandlerImpl {
	return &SentenceHandlerImpl{
		sentenceService: sentenceService,
	}
}
