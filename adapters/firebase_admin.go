package adapters

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func InitializeFirebaseAdmin(credentialsFilePath string) *firebase.App {
	opt := option.WithCredentialsFile(credentialsFilePath)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Println("error initializing firebase admin: ", err)
		os.Exit(1)
	}

	return app
}