package converter

import (
	pb "github.com/gdsc-ys/fluentify-server/gen/proto"
	"github.com/gdsc-ys/fluentify-server/src/model"
)

func ToCompactTopicDTO(topic model.Topic) *pb.CompactTopicDTO {
	return &pb.CompactTopicDTO{
		Id:           topic.Id,
		Title:        topic.Title,
		ThumbnailUrl: topic.ThumbnailUrl,
	}
}

func ToTopicDTO(topic model.Topic) *pb.TopicDTO {
	return &pb.TopicDTO{
		Id:           topic.Id,
		Title:        topic.Title,
		ThumbnailUrl: topic.ThumbnailUrl,
		SentenceIds:  topic.SentenceIds,
		SceneIds:     topic.SceneIds,
	}
}
