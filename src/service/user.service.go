package service

import (
	firebase "firebase.google.com/go/v4"
	"github.com/gdsc-ys/fluentify-server/src/model"
)

type UserService interface {
	GetUser(id string) model.User
}

type UserServiceImpl struct {
	firebaseApp *firebase.App
}

func (service *UserServiceImpl) GetUser(id string) model.User {
	dummyUser := model.User{
		Id:           "fake",
		Name:         "fake",
		Age:          1,
		DisorderType: model.DISORDER_TYPE_HEARING,
	}

	return dummyUser
}

func UserServiceInit(firebaseApp *firebase.App) *UserServiceImpl {
	return &UserServiceImpl{
		firebaseApp: firebaseApp,
	}
}
