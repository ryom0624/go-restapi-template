package store

import (
	"errors"
	"gorm.io/gorm"
	"webapp/auth"
	"webapp/entity"
)

var (
	ErrUserDefinedRecordNotFound = errors.New("ユーザーが見つかりませんでした。")
)

type UserDefinedRecordRepositoryImpl struct {
	db   *DB
	auth auth.Client
}

func NewUserDefinedRecordRepository(db *DB, auth auth.Client) *UserDefinedRecordRepositoryImpl {
	return &UserDefinedRecordRepositoryImpl{db: db, auth: auth}
}

func (drr *UserDefinedRecordRepositoryImpl) GetList(userId string, limit, offset int) ([]entity.UserDefinedRecord, error) {

	var records []entity.UserDefinedRecord
	if err := drr.db.
		Model(&entity.UserDefinedRecord{}).
		Where("user_id = ?", userId).
		Limit(limit).
		Offset(offset).
		Find(&records).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserDefinedRecordNotFound
		}
		return nil, err
	}

	return records, nil
}

func (drr *UserDefinedRecordRepositoryImpl) GetById(userId string, id int64) (*entity.UserDefinedRecord, error) {

	var record entity.UserDefinedRecord
	if err := drr.db.
		Model(&entity.UserDefinedRecord{}).
		Where("user_id = ?", userId).
		Where("id = ?", id).
		Find(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserDefinedRecordNotFound
		}
		return nil, err
	}

	return &record, nil
}

func (drr *UserDefinedRecordRepositoryImpl) Register(userId string, record *entity.UserDefinedRecord) (*entity.UserDefinedRecord, error) {
	tx := drr.db.Begin()
	defer tx.Rollback()
	if err := tx.Create(&record).Error; err != nil {
		return nil, err
	}
	tx.Commit()
	return drr.GetById(userId, record.Id)
}

func (drr *UserDefinedRecordRepositoryImpl) Update(userId string, record *entity.UserDefinedRecord) (*entity.UserDefinedRecord, error) {
	tx := drr.db.Begin()
	defer tx.Rollback()

	if err := tx.Model(&entity.UserDefinedRecord{}).
		Where("user_id = ?", userId).
		Where("id = ?", record.Id).
		Updates(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserDefinedRecordNotFound
		}
		return nil, err
	}
	tx.Commit()

	return drr.GetById(userId, record.Id)
}

func (drr *UserDefinedRecordRepositoryImpl) Delete(userId string, recordId int64) error {
	tx := drr.db.Begin()
	defer tx.Rollback()

	var record entity.UserDefinedRecord
	if err := tx.Model(&entity.UserDefinedRecord{}).
		Where("user_id = ?", userId).
		Where("id = ?", recordId).
		First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserDefinedRecordNotFound
		}
	}

	if err := tx.Model(&entity.UserDefinedRecord{}).
		Where("user_id = ?", userId).
		Where("id = ?", recordId).
		Delete(record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserDefinedRecordNotFound
		}
		return err
	}
	tx.Commit()

	return nil
}
