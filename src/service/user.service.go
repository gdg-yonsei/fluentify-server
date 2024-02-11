package service

import (
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

func UpdateUser(id string, updateUserDTO map[string]interface{}) (model.User, error) {

	dummyUser := model.User{
		Id:           "fake",
		Name:         "fake",
		Age:          1,
		DisorderType: model.DISORDER_TYPE_HEARING,
	}

	for field, value := range updateUserDTO {
		switch field {
		case "name":
			dummyUser.Name = value.(string)
		case "age":
			dummyUser.Age = value.(int)
			if dummyUser.Age < 0 {
				return dummyUser, &model.UserValidationError{Message: "Age must be greater than 0"}
			}
		case "disorderType":
			dummyUser.DisorderType = value.(model.DisorderType)
		}
	}

	return dummyUser, nil
}

func DeleteUser(id string) string {
	return id
}
