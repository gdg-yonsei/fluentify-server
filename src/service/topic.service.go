package service

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/gdsc-ys/fluentify-server/src/model"
	"time"
)

type TopicService interface {
	ListTopics() ([]model.Topic, error)
	GetTopic(id string) (model.Topic, error)
}

const (
	topicCollection = "topic"
	defaultTimeout  = 5 * time.Second
)

type TopicServiceImpl struct {
	firestoreClient *firestore.Client
}

func (service *TopicServiceImpl) ListTopics() ([]model.Topic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	topicDocs, err := service.firestoreClient.Collection(topicCollection).Documents(ctx).GetAll()
	if err != nil || len(topicDocs) == 0 {
		return nil, err
	}

	topicSlice := make([]model.Topic, len(topicDocs))
	for i, doc := range topicDocs {
		var topic = model.Topic{}
		if err := doc.DataTo(&topic); err != nil {
			return nil, err
		}
		topic.Id = doc.Ref.ID
		topicSlice[i] = topic
	}

	return topicSlice, nil
}

func (service *TopicServiceImpl) GetTopic(id string) (model.Topic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	topicDoc, err := service.firestoreClient.
		Collection(topicCollection).
		Where(firestore.DocumentID, "==", id).Limit(1).
		Documents(ctx).
		GetAll()

	if err != nil || len(topicDoc) != 1 {
		return model.Topic{}, err
	}

	singleTopicDoc := topicDoc[0]
	var topic = model.Topic{}
	if err := singleTopicDoc.DataTo(&topic); err != nil {
		return model.Topic{}, err
	}
	topic.Id = singleTopicDoc.Ref.ID

	return topic, nil
}

func TopicServiceInit(firestoreClient *firestore.Client) *TopicServiceImpl {
	return &TopicServiceImpl{
		firestoreClient: firestoreClient,
	}
}
