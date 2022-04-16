package outbound

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App

func InitializeFirebaseAdmin() {
	opt := option.WithCredentialsFile("./pets-hotel-develop-firebase-adminsdk.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal("error initializing firebase admin")
		os.Exit(1)
	}

	FirebaseApp = app
}