package service

import (
	"webapp/entity"
	"webapp/store"
)

var _ UserRepository = (*store.UserRepositoryImpl)(nil)
var _ DailyRecordRepository = (*store.DailyRecordRepositoryImpl)(nil)

//go:generate go run github.com/matryer/moq -out ./moq_test.go . UserRepository DailyRecordRepository
type UserRepository interface {
	GetById(string) (*entity.User, error)
	Register(*entity.User) (*entity.User, error)
	Update(*entity.User) (*entity.User, error)
}

type UserDefinedRecordRepository interface {
	GetList(string, int, int) ([]entity.UserDefinedRecord, error)
	GetById(string, int64) (*entity.UserDefinedRecord, error)
	Register(string, *entity.UserDefinedRecord) (*entity.UserDefinedRecord, error)
	Update(string, *entity.UserDefinedRecord) (*entity.UserDefinedRecord, error)
	Delete(string, int64) error
}

type DailyRecordRepository interface {
	GetList(string, int, int) ([]entity.DailyRecord, error)
	GetById(string, int64) (*entity.DailyRecord, error)
	Register(string, *entity.DailyRecord) (*entity.DailyRecord, error)
	Update(string, *entity.DailyRecord) (*entity.DailyRecord, error)
	Delete(string, int64) error
}
>>>>>>> Stashed changes
