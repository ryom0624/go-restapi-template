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

var _ DailyRecordHandler = (*service.DailyRecordService)(nil)
var _ DailyRecordGetter = (*service.DailyRecordService)(nil)
var _ DailyRecordUpdater = (*service.DailyRecordService)(nil)
var _ DailyRecordDeleter = (*service.DailyRecordService)(nil)

//go:generate go run github.com/matryer/moq -out ./moq_test.go . DailyRecordHandler DailyRecordRegister DailyRecordGetter DailyRecordUpdater DailyRecordDeleter
type DailyRecordHandler interface {
	DailyRecordRegister
	DailyRecordGetter
	DailyRecordUpdater
	DailyRecordDeleter
}

type DailyRecordRegister interface {
	RegisterDailyRecord(ctx context.Context, record *entity.DailyRecord) (*entity.DailyRecord, error)
}

type DailyRecordGetter interface {
	GetDailyRecordList(ctx context.Context, limit, offset int) ([]entity.DailyRecord, error)
	GetDailyRecordById(ctx context.Context, recordId int64) (*entity.DailyRecord, error)
}

type DailyRecordUpdater interface {
	UpdateDailyRecord(ctx context.Context, user *entity.DailyRecord) (*entity.DailyRecord, error)
}

type DailyRecordDeleter interface {
	DeleteDailyRecord(ctx context.Context, recordId int64) error
}

var _ UserDefinedRecordHandler = (*service.UserDefinedRecordService)(nil)
var _ UserDefinedRecordGetter = (*service.UserDefinedRecordService)(nil)
var _ UserDefinedRecordUpdater = (*service.UserDefinedRecordService)(nil)
var _ UserDefinedRecordDeleter = (*service.UserDefinedRecordService)(nil)

//go:generate go run github.com/matryer/moq -out ./moq_test.go . UserDefinedRecordHandler UserDefinedRecordRegister UserDefinedRecordGetter UserDefinedRecordUpdater UserDefinedRecordDeleter
type UserDefinedRecordHandler interface {
	UserDefinedRecordRegister
	UserDefinedRecordGetter
	UserDefinedRecordUpdater
	UserDefinedRecordDeleter
}

type UserDefinedRecordRegister interface {
	RegisterUserDefinedRecord(ctx context.Context, record *entity.UserDefinedRecord) (*entity.UserDefinedRecord, error)
}

type UserDefinedRecordGetter interface {
	GetUserDefinedRecordList(ctx context.Context, limit int, offset int) ([]entity.UserDefinedRecord, error)
	GetUserDefinedRecordById(ctx context.Context, id int64) (*entity.UserDefinedRecord, error)
}

type UserDefinedRecordUpdater interface {
	UpdateUserDefinedRecord(ctx context.Context, record *entity.UserDefinedRecord) (*entity.UserDefinedRecord, error)
}

type UserDefinedRecordDeleter interface {
	DeleteUserDefinedRecord(ctx context.Context, id int64) error
}
>>>>>>> Stashed changes
