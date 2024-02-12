package converter

import (
	pb "github.com/gdsc-ys/fluentify-server/gen/idl/proto"
	"github.com/gdsc-ys/fluentify-server/src/model"
)

func ConvertUser(user model.User) pb.UserDTO {
	return pb.UserDTO{
		Id:           user.Id,
		Name:         user.Name,
		Age:          int32(user.Age),
		DisorderType: convertDisorderType(user.DisorderType),
	}
}

func convertDisorderType(disorderType model.DisorderType) pb.DisorderType {
	switch disorderType {
	default:
		return pb.DisorderType_DISORDER_TYPE_UNSPECIFIED
	}
}

func ConvertDisorderType(disorderType pb.DisorderType) model.DisorderType {
	switch disorderType {
	default:
		return model.DISORDER_TYPE_SIGHT
	}
}
