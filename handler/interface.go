package handler

import (
	"context"
	"webapp/entity"
	"webapp/service"
)

var _ UserHandler = (*service.UserService)(nil)
var _ UserGetter = (*service.UserService)(nil)
var _ UserUpdater = (*service.UserService)(nil)

//go:generate go run github.com/matryer/moq -out ./moq_test.go . UserHandler UserRegister UserGetter UserUpdater
type UserHandler interface {
	UserRegister
	UserGetter
	UserUpdater
}

type UserRegister interface {
	RegisterUser(ctx context.Context, user *entity.User) (*entity.User, error)
}

type UserGetter interface {
	GetUserById(ctx context.Context) (*entity.User, error)
}

type UserUpdater interface {
	UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error)
}
