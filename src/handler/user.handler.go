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
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
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

	user, err := handler.userService.GetUser(request.Id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "invalid id")
	}
	userDTO := converter.ConvertUser(user)

	return c.JSON(http.StatusOK, pb.GetUserResponse{User: &userDTO})
}

func (handler *UserHandlerImpl) UpdateUser(c echo.Context) error {
	var request = pb.UpdateUserRequest{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if request.Id == "" {
		return c.JSON(http.StatusBadRequest, "Id is required")
	}

	userUpdateDTO := make(map[string]interface{})

	switch {
	case request.GetName() != "":
		userUpdateDTO["name"] = request.GetName()

	case request.GetAge() != 0:
		userUpdateDTO["age"] = int(request.GetAge())

	case request.GetDisorderType() != 0:
		userUpdateDTO["disorderType"] = converter.ConvertDisorderType(request.GetDisorderType())

	default:
		return c.JSON(http.StatusBadRequest, "At least one field is required")
	}

	if user, err := handler.userService.UpdateUser(request.Id, userUpdateDTO); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	} else {
		userDTO := converter.ConvertUser(user)
		return c.JSON(http.StatusOK, pb.UpdateUserResponse{User: &userDTO})
	}
}

func (handler *UserHandlerImpl) DeleteUser(c echo.Context) error {
	var request = pb.DeleteUserRequest{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if request.Id == "" {
		return c.JSON(http.StatusBadRequest, "Id is required")
	}

	deletedUserId := handler.userService.DeleteUser(request.Id)

	return c.JSON(http.StatusOK, pb.DeleteUserResponse{Id: deletedUserId})
}

func UserHandlerInit(userService service.UserService) *UserHandlerImpl {
	return &UserHandlerImpl{
		userService: userService,
	}
}
