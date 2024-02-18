package service

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/gdsc-ys/fluentify-server/src/constant"
	"github.com/gdsc-ys/fluentify-server/src/model"
)

type SentenceService interface {
	GetSentence(id string) (model.Sentence, error)
}

type SentenceServiceImpl struct {
	firestoreClient *firestore.Client
}

func (service *SentenceServiceImpl) GetSentence(id string) (model.Sentence, error) {
	ctx, cancel := context.WithTimeout(context.Background(), constant.FirestoreDefaultTimeout)
	defer cancel()

	sentenceDoc, err := service.firestoreClient.
		Collection(constant.SentenceCollection).
		Doc(id).
		Get(ctx)

	if err != nil {
		return model.Sentence{}, err
	}

	var sentence = model.Sentence{}
	if err := sentenceDoc.DataTo(&sentence); err != nil {
		return model.Sentence{}, err
	}
	sentence.Id = sentenceDoc.Ref.ID

	return sentence, nil
}

func SentenceServiceInit(firestoreClient *firestore.Client) *SentenceServiceImpl {
	return &SentenceServiceImpl{
		firestoreClient: firestoreClient,
	}
}
