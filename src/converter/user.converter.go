package converter

import (
	pb "github.com/gdsc-ys/fluentify-server/gen/idl/proto"
	"github.com/gdsc-ys/fluentify-server/src/domain"
)

func ConvertUser(user domain.User) pb.UserDTO {
	return pb.UserDTO{
		Id:           user.Id,
		Name:         user.Name,
		Age:          int32(user.Age),
		DisorderType: convertDisorderType(user.DisorderType),
	}
}

func convertDisorderType(disorderType domain.DisorderType) pb.DisorderType {
	switch disorderType {
	default:
		return pb.DisorderType_DISORDER_TYPE_UNSPECIFIED
	}
}
