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

	userUpdateDTO := map[string]interface{}{}

	if name := request.GetName(); name != "" {
		userUpdateDTO["name"] = name
	}
	if age := request.GetAge(); age != 0 {
		userUpdateDTO["age"] = int(age)
	}
	if disorderType := request.GetDisorderType(); disorderType != 0 {
		userUpdateDTO["disorderType"] = disorderType.Number()
	}

	if len(userUpdateDTO) == 0 {
		return c.JSON(http.StatusBadRequest, "At least one field is required")
	}

	userUpdateDTO["uid"] = request.GetId()

	if user, err := handler.userService.UpdateUser(userUpdateDTO); err != nil {
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

	err := handler.userService.DeleteUser(request.Id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "invalid id")
	}

	return c.JSON(http.StatusOK, pb.DeleteUserResponse{Id: request.Id})
}

func UserHandlerInit(userService service.UserService) *UserHandlerImpl {
	return &UserHandlerImpl{
		userService: userService,
	}
}
