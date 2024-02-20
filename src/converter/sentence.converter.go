package converter

import (
	pb "github.com/gdsc-ys/fluentify-server/gen/proto"
	"github.com/gdsc-ys/fluentify-server/src/model"
)

func ToSentenceDTO(sentence model.Sentence, exampleVideoUrl string) *pb.SentenceDTO {
	return &pb.SentenceDTO{
		Id:              sentence.Id,
		Text:            sentence.Text,
		ExampleVideoUrl: exampleVideoUrl,
	}
}
