package config

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"log"
)

func InitializeFirebaseApp() *firebase.App {
	firebaseApp, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v\n", err)
	}

	return firebaseApp
}
