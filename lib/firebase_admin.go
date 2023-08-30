package lib

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

type Client struct {
	*auth.Client
}

func NewClient(ctx context.Context, env string, projectId string, keyfileJSON []byte) (*Client, error) {
	// todo: keyfile is required when testing on ci

	app, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: projectId,
	})
	if err != nil {
		return nil, err
	}
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	return &Client{client}, nil
}
