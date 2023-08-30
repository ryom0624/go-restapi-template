package store

import (
	"errors"
	"gorm.io/gorm"
	"webapp/auth"
	"webapp/entity"
)

var (
	ErrUserNotFound = errors.New("ユーザーが見つかりませんでした。")
)

type UserRepositoryImpl struct {
	db   *DB
	auth auth.Client
}

func NewUserRepository(db *DB, auth auth.Client) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db, auth: auth}
}

func (ur *UserRepositoryImpl) GetById(id string) (*entity.User, error) {

	var user entity.User
	if err := ur.db.
		Model(&entity.User{}).
		Where("user.id = ?", id).
		Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepositoryImpl) Register(user *entity.User) (*entity.User, error) {
	tx := ur.db.Begin()
	defer tx.Rollback()
	if err := tx.Create(&user).Error; err != nil {
		return nil, err
	}
	tx.Commit()
	return ur.GetById(user.Id)
}

func (ur *UserRepositoryImpl) Update(user *entity.User) (*entity.User, error) {
	tx := ur.db.Begin()
	defer tx.Rollback()

	if err := tx.Model(user).Where("id = ?", user.Id).Updates(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	tx.Commit()

	return ur.GetById(user.Id)
}
