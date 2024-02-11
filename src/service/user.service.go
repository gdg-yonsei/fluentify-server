package service

import (
	"firebase.google.com/go/v4/auth"
	"github.com/gdsc-ys/fluentify-server/src/model"
)

type UserService interface {
	GetUser(id string) model.User
}

type UserServiceImpl struct {
	authClient *auth.Client
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

func UserServiceInit(authClient *auth.Client) *UserServiceImpl {
	return &UserServiceImpl{
		authClient: authClient,
	}
}
