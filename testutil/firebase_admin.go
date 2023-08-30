package testutil

import (
	"context"
	"firebase.google.com/go/auth"
	"webapp/config"
	"webapp/lib"
)

func SetupFirebaseAuthentication(email, password string) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	ctx := context.Background()

	cli, err := lib.NewClient(ctx, "localhost", cfg.GoogleProjectId, cfg.FirebaseCredentialJSON)
	if err != nil {
		return err
	}

	params := (&auth.UserToCreate{}).
		Email(email).
		EmailVerified(false).
		Password(password).
		Disabled(false)

	_, err = cli.CreateUser(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

// generate id token for test
func GenerateIdToken(email, pass string) (string, error) {
	cfg, err := config.New()
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	cli, err := lib.NewClient(ctx, "localhost", cfg.GoogleProjectId, cfg.FirebaseCredentialJSON)
	if err != nil {
		return "", err
	}

	token, err := cli.CustomToken(ctx, email)
	if err != nil {
		return "", err
	}

	idToken, err := cli.VerifyIDToken(ctx, token)
	if err != nil {
		return "", err
	}

	return idToken.UID, nil
}

func TeardownFirebaseAuthentication(email string) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	ctx := context.Background()

	cli, err := lib.NewClient(ctx, "localhost", cfg.GoogleProjectId, cfg.FirebaseCredentialJSON)
	if err != nil {
		return err
	}

	user, err := cli.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	err = cli.DeleteUser(ctx, user.UID)
	if err != nil {
		return err
	}
	return nil
}
