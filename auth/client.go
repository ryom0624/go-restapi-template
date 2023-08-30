package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"webapp/lib"
)

type Client interface {
	FillContext(r *http.Request) (*http.Request, error)
}

type NoopClient struct {
	userId string
	email  string
}

func NewNoopClient(userId, email string) (NoopClient, error) {
	return NoopClient{
		userId: userId,
		email:  email,
	}, nil
}

func (c NoopClient) FillContext(r *http.Request) (*http.Request, error) {
	ctx := r.Context()
	ctx = SetUserID(ctx, c.userId)
	ctx = SetUserEmail(ctx, c.email)
	clone := r.Clone(ctx)
	return clone, nil
}

type FirebaseAdmin struct {
	cli *lib.Client
}

type Token struct {
	UID    string                 `json:"uid"`
	Email  string                 `json:"email"`
	Claims map[string]interface{} `json:"-"`
}

func NewFirebaseAdminClient(ctx context.Context, env, googleProjectId string, keyfileJSON []byte) (FirebaseAdmin, error) {
	c, err := lib.NewClient(ctx, env, googleProjectId, keyfileJSON)
	if err != nil {
		return FirebaseAdmin{}, err
	}
	return FirebaseAdmin{cli: c}, nil
}

func (c FirebaseAdmin) FillContext(r *http.Request) (*http.Request, error) {
	ctx := r.Context()
	token, err := c.verifyIDToken(ctx, r)
	if err != nil {
		return nil, err
	}
	if token.UID == "" {
		return nil, errors.New("uid not found in token")
	}
	ctx = SetUserID(ctx, token.UID)
	ctx = SetUserEmail(ctx, token.Email)
	clone := r.Clone(ctx)
	return clone, nil
}

func (c FirebaseAdmin) verifyIDToken(ctx context.Context, r *http.Request) (*Token, error) {
	header := r.Header.Get("Authorization")
	//log.Printf("header: %s", header)
	bearer := strings.Split(header, "Bearer ")
	if len(bearer) != 2 {
		return nil, errors.New("invalid authorization header")
	}
	//log.Printf("bearer: %s", bearer[1])
	t, err := c.cli.VerifyIDToken(ctx, bearer[1])
	if err != nil {
		return nil, fmt.Errorf("failed to verify id token: %w", err)
	}
	userInfo, err := c.cli.GetUser(ctx, t.UID)
	if err != nil {
		return nil, err
	}
	return &Token{
		UID:    t.UID,
		Email:  userInfo.Email,
		Claims: t.Claims,
	}, nil
}

func (c FirebaseAdmin) DeleteUser(ctx context.Context, id string) error {
	err := c.cli.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

type UserId struct{}

func GetUserID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(UserId{}).(string)
	return id, ok
}

func SetUserID(ctx context.Context, uid string) context.Context {
	return context.WithValue(ctx, UserId{}, uid)
}

type UserEmail struct{}

func GetUserEmail(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(UserEmail{}).(string)
	return id, ok
}

func SetUserEmail(ctx context.Context, uid string) context.Context {
	return context.WithValue(ctx, UserEmail{}, uid)
}
