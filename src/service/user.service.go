package service

import (
	"context"

	"firebase.google.com/go/v4/auth"
	"github.com/gdsc-ys/fluentify-server/src/model"
)

func GetUser(id string) model.User {
	dummyUser := model.User{
		Id:           "fake",
		Name:         "fake",
		Age:          1,
		DisorderType: model.DISORDER_TYPE_HEARING,
	}

	return dummyUser
}

func UpdateUser(client *auth.Client, updateUserDTO map[string]interface{}) (model.User, error) {

	ctx := context.Background()
	params := (&auth.UserToUpdate{})
	customClaims := make(map[string]interface{})

	for field, value := range updateUserDTO {
		switch field {
		case "name":
			params = params.DisplayName(value.(string))
		case "age":
			if value.(int) < 0 {
				return model.User{}, &model.UserValidationError{Message: "Age must be greater than 0"}
			}
			customClaims["age"] = value.(int)
		case "disorderType":
			customClaims["disorderType"] = value.(string)
		}
	}
	params = params.CustomClaims(customClaims)

	userRecord, err := client.UpdateUser(ctx, updateUserDTO["uid"].(string), params)
	if err != nil {
		return model.User{}, err
	}

	user := convertRecordToUser(userRecord)
	return user, nil

}

func convertRecordToUser(record *auth.UserRecord) model.User {
	user := model.User{
		Id:           record.UID,
		Name:         record.DisplayName,
		Age:          record.CustomClaims["age"].(int),
		DisorderType: model.DISORDER_TYPE_HEARING,
	}

	return user
}
