package handler

import (
	"net/http"

	pb "github.com/gdsc-ys/fluentify-server/gen/idl/proto"
	"github.com/gdsc-ys/fluentify-server/src/converter"
	"github.com/gdsc-ys/fluentify-server/src/service"
	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	GetUser(c echo.Context) error
}

type UserHandlerImpl struct {
	userService service.UserService
}

func (handler *UserHandlerImpl) GetUser(c echo.Context) error {
	var request = pb.GetUserRequest{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if request.Id == "" {
		return c.JSON(http.StatusBadRequest, "Id is required")
	}

	user := handler.userService.GetUser(request.Id)
	userDTO := converter.ConvertUser(user)

	return c.JSON(http.StatusOK, pb.GetUserResponse{User: &userDTO})
}

func UserHandlerInit(userService service.UserService) *UserHandlerImpl {
	return &UserHandlerImpl{
		userService: userService,
	}
}
