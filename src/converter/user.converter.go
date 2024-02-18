package converter

import (
	pb "github.com/gdsc-ys/fluentify-server/gen/proto"
	"github.com/gdsc-ys/fluentify-server/src/model"
)

func ToUserDTO(user model.User) *pb.UserDTO {
	return &pb.UserDTO{
		Id:   user.Id,
		Name: user.Name,
		Age:  int32(user.Age),
	}
}
