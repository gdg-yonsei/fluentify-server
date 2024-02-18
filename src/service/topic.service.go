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
	topicCollection    = "topic"
	sentenceCollection = "sentence"
	sceneCollection    = "scene"
	topicIdField       = "topic_id"
	defaultTimeout     = 5 * time.Second
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
		Doc(id).
		Get(ctx)

	if err != nil {
		return model.Topic{}, err
	}

	var topic = model.Topic{}
	if err := topicDoc.DataTo(&topic); err != nil {
		return model.Topic{}, err
	}
	topic.Id = topicDoc.Ref.ID

	sentenceIds, err := service.getSentenceIdsByTopicId(ctx, id)
	if err != nil {
		return model.Topic{}, err
	}
	topic.SentenceIds = sentenceIds

	sceneIds, err := service.getSceneIdsByTopicId(ctx, id)
	if err != nil {
		return model.Topic{}, err
	}
	topic.SceneIds = sceneIds

	return topic, nil
}

func (service *TopicServiceImpl) getSentenceIdsByTopicId(ctx context.Context, topicId string) ([]string, error) {
	sentenceDocs, err := service.firestoreClient.Collection(sentenceCollection).Where(topicIdField, "==", topicId).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	sentenceIds := make([]string, len(sentenceDocs))
	for i, doc := range sentenceDocs {
		sentenceIds[i] = doc.Ref.ID
	}

	return sentenceIds, nil
}

func (service *TopicServiceImpl) getSceneIdsByTopicId(ctx context.Context, topicId string) ([]string, error) {
	sceneDocs, err := service.firestoreClient.Collection(sceneCollection).Where(topicIdField, "==", topicId).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	sceneIds := make([]string, len(sceneDocs))
	for i, doc := range sceneDocs {
		sceneIds[i] = doc.Ref.ID
	}

	return sceneIds, nil
}

func TopicServiceInit(firestoreClient *firestore.Client) *TopicServiceImpl {
	return &TopicServiceImpl{
		firestoreClient: firestoreClient,
	}
}
