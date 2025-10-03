package firebaseadmin

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

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

	return authClient
}