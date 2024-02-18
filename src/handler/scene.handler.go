package handler

import (
	pb "github.com/gdsc-ys/fluentify-server/gen/proto"
	"github.com/gdsc-ys/fluentify-server/src/converter"
	"github.com/gdsc-ys/fluentify-server/src/model"
	"github.com/gdsc-ys/fluentify-server/src/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type SceneHandler interface {
	GetScene(c echo.Context) error
}

type SceneHandlerImpl struct {
	sceneService   service.SceneService
	storageService service.StorageService
}

func (handler *SceneHandlerImpl) GetScene(c echo.Context) error {
	var request = pb.GetSceneRequest{}
	if err := c.Bind(&request); err != nil {
		return model.NewCustomHTTPError(http.StatusBadRequest, err)
	}
	if request.GetId() == "" {
		return model.NewCustomHTTPError(http.StatusBadRequest, "scene id is required")
	}

	scene, err := handler.sceneService.GetScene(request.GetId())
	if err != nil {
		return model.NewCustomHTTPError(http.StatusInternalServerError, err)
	}

	thumbnailURL, err := handler.storageService.GetFileUrl(scene.ImageUrl)
	if err != nil {
		return model.NewCustomHTTPError(http.StatusInternalServerError, err)
	}
	scene.ImageUrl = thumbnailURL

	return c.JSON(http.StatusOK, converter.ToSceneDTO(scene))
}

func SceneHandlerInit(sceneService service.SceneService, storageService service.StorageService) *SceneHandlerImpl {
	return &SceneHandlerImpl{
		sceneService:   sceneService,
		storageService: storageService,
	}
}
