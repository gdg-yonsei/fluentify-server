package handler

import (
	"net/http"

	pb "github.com/gdsc-ys/fluentify-server/gen/idl/proto"
	"github.com/gdsc-ys/fluentify-server/src/converter"
	"github.com/gdsc-ys/fluentify-server/src/service"
	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	var request = pb.GetUserRequest{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if request.Id == "" {
		return c.JSON(http.StatusBadRequest, "Id is required")
	}

	user := service.GetUser(request.Id)
	userDTO := converter.ConvertUser(user)

	return c.JSON(http.StatusOK, pb.GetUserResponse{User: &userDTO})
}

func UpdateUser(c echo.Context) error {
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

	user := service.UpdateUser(request.Id, userUpdateDTO)
	userDTO := converter.ConvertUser(user)

	return c.JSON(http.StatusOK, pb.UpdateUserResponse{User: &userDTO})

}

func DeleteUser(c echo.Context) error {
	var request = pb.DeleteUserRequest{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if request.Id == "" {
		return c.JSON(http.StatusBadRequest, "Id is required")
	}

	deletedUserId := service.DeleteUser(request.Id)

	return c.JSON(http.StatusOK, pb.DeleteUserResponse{Id: deletedUserId})
}
