package auth

import (
	"fmt"
	"context"
	"log"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

var messagingClient *messaging.Client

func InitFirebase(serviceAccountPath string) error {
    opt := option.WithCredentialsFile(serviceAccountPath)
    app, err := firebase.NewApp(context.Background(), nil, opt)
    if err != nil {
        return err
    }
    
    messagingClient, err = app.Messaging(context.Background())
    if err != nil {
        return err
    }
    
    log.Println("Firebase initialized successfully")
    return nil
}

func GetMessagingClient() *messaging.Client {
    return messagingClient
}

func NewAuthClient(c context.Context, serviceAccountPath string) (*auth.Client, error) {
	var (
		app *firebase.App
		err error
	)
	app, err = firebase.NewApp(c, nil, option.WithCredentialsFile(serviceAccountPath))
	if err != nil {
		return nil, fmt.Errorf("failed to create Firebase app: %v", err)
	}

	authClient, err := app.Auth(c)
	if err != nil {
		return nil, fmt.Errorf("failed to create Firebase auth client: %v", err)
	}

	return authClient, err
}