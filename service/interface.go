package service

import (
	"webapp/entity"
	"webapp/store"
)

var _ UserRepository = (*store.UserRepositoryImpl)(nil)

//go:generate go run github.com/matryer/moq -out ./moq_test.go . UserRepository DailyRecordRepository
type UserRepository interface {
	GetById(string) (*entity.User, error)
	Register(*entity.User) (*entity.User, error)
	Update(*entity.User) (*entity.User, error)
}
