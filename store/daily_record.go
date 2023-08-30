package store

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"webapp/auth"
	"webapp/entity"
)

var (
	ErrDailyRecordNotFound = errors.New("該当のレコードが見つかりませんでした。")
)

type DailyRecordRepositoryImpl struct {
	db   *DB
	auth auth.Client
}

func NewDailyRecordRepository(db *DB, auth auth.Client) *DailyRecordRepositoryImpl {
	return &DailyRecordRepositoryImpl{db: db, auth: auth}
}

func (drr *DailyRecordRepositoryImpl) GetList(userId string, limit, offset int) ([]entity.DailyRecord, error) {

	var records []entity.DailyRecord
	if err := drr.db.
		Model(&entity.DailyRecord{}).
		Preload("OptionalRecords").
		Where("user_id = ?", userId).
		Limit(limit).
		Offset(offset).
		Find(&records).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrDailyRecordNotFound
		}
		return nil, err
	}

	return records, nil
}

func (drr *DailyRecordRepositoryImpl) GetById(userId string, id int64) (*entity.DailyRecord, error) {

	var record entity.DailyRecord
	if err := drr.db.
		Preload("OptionalRecords").
		Model(&entity.DailyRecord{}).
		Where("user_id = ?", userId).
		Where("id = ?", id).
		First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrDailyRecordNotFound
		}
		return nil, err
	}

	return &record, nil
}

func (drr *DailyRecordRepositoryImpl) Register(userId string, record *entity.DailyRecord) (*entity.DailyRecord, error) {
	tx := drr.db.Begin()
	defer tx.Rollback()

	var userDefinedIds []int64
	for i := range record.OptionalRecords {
		userDefinedIds = append(userDefinedIds, record.OptionalRecords[i].UserDefinedId)
	}

	var userDefinedRecords []entity.UserDefinedRecord
	if err := tx.Model(&entity.UserDefinedRecord{}).
		Where("user_id = ?", userId).
		Where("id IN ?", userDefinedIds).
		Find(&userDefinedRecords).Error; err != nil {
		return nil, err
	}
	var validateUserDefinedIdsMap = make(map[int64]bool)
	for i := range userDefinedRecords {
		validateUserDefinedIdsMap[userDefinedRecords[i].Id] = true
	}

	for i := range record.OptionalRecords {
		if _, ok := validateUserDefinedIdsMap[record.OptionalRecords[i].UserDefinedId]; !ok {
			return nil, fmt.Errorf("user_defined_id %d is not found", record.OptionalRecords[i].UserDefinedId)
		}
	}

	if err := tx.Create(&record).Error; err != nil {
		return nil, err
	}

	tx.Commit()
	return drr.GetById(userId, record.Id)
}

func (drr *DailyRecordRepositoryImpl) Update(userId string, record *entity.DailyRecord) (*entity.DailyRecord, error) {
	tx := drr.db.Begin()
	defer tx.Rollback()

	if len(record.OptionalRecords) > 0 {
		// validate user_defined record
		var userDefinedIds []int64
		for i := range record.OptionalRecords {
			userDefinedIds = append(userDefinedIds, record.OptionalRecords[i].UserDefinedId)
		}

		var userDefinedRecords []entity.UserDefinedRecord
		if err := tx.Model(&entity.UserDefinedRecord{}).
			Where("user_id = ?", userId).
			Where("id IN ?", userDefinedIds).
			Find(&userDefinedRecords).Error; err != nil {
			return nil, err
		}
		var validateUserDefinedIdsMap = make(map[int64]bool)
		for i := range userDefinedRecords {
			validateUserDefinedIdsMap[userDefinedRecords[i].Id] = true
		}

		for i := range record.OptionalRecords {
			if _, ok := validateUserDefinedIdsMap[record.OptionalRecords[i].UserDefinedId]; !ok {
				return nil, fmt.Errorf("user_defined_id %d is not found", record.OptionalRecords[i].UserDefinedId)
			}
		}

		var deleteRecords []entity.OptionalRecord
		if err := tx.
			Select("optional_record.*").
			InnerJoins("INNER JOIN daily_record ON daily_record.id = optional_record.daily_record_id").
			Where("daily_record.user_id = ?", userId).
			Where("daily_record.id = ?", record.Id).
			Find(&deleteRecords).Error; err != nil {
			return nil, err
		}

		if len(deleteRecords) == 0 {
			return nil, fmt.Errorf("daily_record_id %d is not found", record.Id)
		}

		//	update optional record
		if err := tx.
			Delete(deleteRecords).Error; err != nil {
			return nil, err
		}

		newOptionalRecord := make([]entity.OptionalRecord, len(record.OptionalRecords))
		for i := range record.OptionalRecords {
			newOptionalRecord[i] = entity.OptionalRecord{
				DailyRecordId: record.Id,
				UserDefinedId: record.OptionalRecords[i].UserDefinedId,
				Value:         record.OptionalRecords[i].Value,
			}
		}
		if err := tx.Create(&newOptionalRecord).Error; err != nil {
			return nil, err
		}
	}

	// update daily record
	if err := tx.Model(&entity.DailyRecord{}).
		Where("user_id = ?", userId).
		Where("id = ?", record.Id).
		Updates(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrDailyRecordNotFound
		}
		return nil, err
	}
	tx.Commit()

	return drr.GetById(userId, record.Id)
}

func (drr *DailyRecordRepositoryImpl) Delete(userId string, recordId int64) error {
	tx := drr.db.Begin()
	defer tx.Rollback()

	if err := tx.Model(&entity.DailyRecord{}).
		Where("user_id = ?", userId).
		//Where("id = ?", recordId).
		Delete(recordId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrDailyRecordNotFound
		}
		return err
	}
	tx.Commit()

	return nil
}
