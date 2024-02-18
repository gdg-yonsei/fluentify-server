package converter

import (
	pb "github.com/gdsc-ys/fluentify-server/gen/proto"
	"github.com/gdsc-ys/fluentify-server/src/model"
)

func ToSceneDTO(scene model.Scene) *pb.SceneDTO {
	return &pb.SceneDTO{
		Id:       scene.Id,
		Question: scene.Question,
		ImageUrl: scene.ImageUrl,
	}
}
