package handler

import (
	"context"
	"log"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	pb "github.com/gdsc-ys/fluentify-server/gen/idl/proto"
	"github.com/gdsc-ys/fluentify-server/src/converter"
	"github.com/gdsc-ys/fluentify-server/src/service"
	"github.com/labstack/echo/v4"
)

// DI로 뺐으므로 동작 확인을 위한 임의작성합니다
func getAuthClient() *auth.Client {
	ctx := context.Background()

	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v", err)
	}

	return client
}

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
	client := getAuthClient()

	var request = pb.UpdateUserRequest{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// if request.Id == "" {
	// 	return c.JSON(http.StatusBadRequest, "Id is required")
	// }

	userUpdateDTO := make(map[string]interface{})

	userUpdateDTO["uid"] = c.Get("uid")

	switch {
	case request.GetName() != "":
		userUpdateDTO["name"] = request.GetName()

	case request.GetAge() != 0:
		userUpdateDTO["age"] = int(request.GetAge())

	case request.GetDisorderType() != 0:
		userUpdateDTO["disorderType"] = request.GetDisorderType().Number()

	default:
		return c.JSON(http.StatusBadRequest, "At least one field is required")
	}

	if user, err := service.UpdateUser(client, userUpdateDTO); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	} else {
		userDTO := converter.ConvertUser(user)
		return c.JSON(http.StatusOK, pb.UpdateUserResponse{User: &userDTO})
	}
}
