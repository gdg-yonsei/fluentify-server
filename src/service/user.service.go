package service

import (
	"github.com/gdsc-ys/fluentify-server/src/domain"
)

func GetUser(id string) domain.User {
	dummyUser := domain.User{
		Id:           "fake",
		Name:         "fake",
		Age:          1,
		DisorderType: domain.DISORDER_TYPE_UNSPECIFIED,
	}

	return dummyUser
}
