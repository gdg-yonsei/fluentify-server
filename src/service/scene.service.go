package service

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/gdsc-ys/fluentify-server/src/constant"
	"github.com/gdsc-ys/fluentify-server/src/model"
)

type SceneService interface {
	GetScene(id string) (model.Scene, error)
}

type SceneServiceImpl struct {
	firestoreClient *firestore.Client
}

func (service *SceneServiceImpl) GetScene(id string) (model.Scene, error) {
	ctx, cancel := context.WithTimeout(context.Background(), constant.FirestoreDefaultTimeout)
	defer cancel()

	sceneDoc, err := service.firestoreClient.
		Collection(constant.SceneCollection).
		Doc(id).
		Get(ctx)

	if err != nil {
		return model.Scene{}, err
	}

	var scene = model.Scene{}
	if err := sceneDoc.DataTo(&scene); err != nil {
		return model.Scene{}, err
	}
	scene.Id = sceneDoc.Ref.ID

	return scene, nil
}

func SceneServiceInit(firestoreClient *firestore.Client) *SceneServiceImpl {
	return &SceneServiceImpl{
		firestoreClient: firestoreClient,
	}
}
