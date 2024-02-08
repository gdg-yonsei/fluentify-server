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

func UpdateUser(id string, args map[string]interface{}) model.User {

	dummyUser := model.User{
		Id:           "fake",
		Name:         "fake",
		Age:          1,
		DisorderType: model.DISORDER_TYPE_HEARING,
	}

	for field, value := range args {
		switch field {
		case "name":
			dummyUser.Name = value.(string)
		case "age":
			dummyUser.Age = value.(int)
		case "disorderType":
			dummyUser.DisorderType = value.(model.DisorderType)
		}
	}

	return dummyUser
}

func DeleteUser(id string) string {
	return id
}
