package config

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/storage"
	"log"
	"os"
)

func InitializeFirebaseApp() *firebase.App {
	defaultBucketName := os.Getenv("DEFAULT_STORAGE_BUCKET")
	config := &firebase.Config{
		StorageBucket: defaultBucketName + ".appspot.com",
	}

	firebaseApp, err := firebase.NewApp(context.Background(), config)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v\n", err)
	}

	return firebaseApp
}

func NewFirebaseAuthClient(app *firebase.App) *auth.Client {
	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting firebase auth client: %v", err)
	}

	return authClient
}

func NewFirebaseStorageClient(app *firebase.App) *storage.Client {
	storageClient, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalf("error getting firebase storage client: %v", err)
	}

	return storageClient
}

func NewFireStoreClient(app *firebase.App) *firestore.Client {
	firestoreClient, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("error getting firestore client: %v", err)
	}

	return firestoreClient
}
