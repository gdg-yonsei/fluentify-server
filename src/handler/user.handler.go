package handler

import (
	"github.com/gdsc-ys/fluentify-server/src/model"
	"net/http"

	pb "github.com/gdsc-ys/fluentify-server/gen/proto"
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
		return model.NewCustomHTTPError(http.StatusBadRequest, err)
	}

	if request.Id == "" {
		return model.NewCustomHTTPError(http.StatusBadRequest, "id is required")
	}

	user, err := handler.userService.GetUser(request.Id)
	if err != nil {
		return err
	}
	userDTO := converter.ToUserDTO(user)

	return c.JSON(http.StatusOK, pb.GetUserResponse{User: userDTO})
}

func (handler *UserHandlerImpl) UpdateUser(c echo.Context) error {
	var request = pb.UpdateUserRequest{}
	if err := c.Bind(&request); err != nil {
		return model.NewCustomHTTPError(http.StatusBadRequest, err)
	}

	userUpdateDTO := map[string]interface{}{}

	if name := request.GetName(); name != "" {
		userUpdateDTO["name"] = name
	}
	if age := request.GetAge(); age != 0 {
		userUpdateDTO["age"] = int(age)
	}

	if len(userUpdateDTO) == 0 {
		return model.NewCustomHTTPError(http.StatusBadRequest, "at least one field is required")
	}

	userUpdateDTO["uid"] = request.GetId()
	user, err := handler.userService.UpdateUser(userUpdateDTO)
	if err != nil {
		return err
	}

	userDTO := converter.ToUserDTO(user)
	return c.JSON(http.StatusOK, pb.UpdateUserResponse{User: userDTO})
}

func (handler *UserHandlerImpl) DeleteUser(c echo.Context) error {
	var request = pb.DeleteUserRequest{}
	if err := c.Bind(&request); err != nil {
		return model.NewCustomHTTPError(http.StatusBadRequest, err)
	}

	if request.Id == "" {
		return model.NewCustomHTTPError(http.StatusBadRequest, "id is required")
	}

	err := handler.userService.DeleteUser(request.Id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, pb.DeleteUserResponse{Id: request.Id})
}

func UserHandlerInit(userService service.UserService) *UserHandlerImpl {
	return &UserHandlerImpl{
		userService: userService,
	}
}
