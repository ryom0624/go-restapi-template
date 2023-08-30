package service

import (
	"context"
	"errors"
	"webapp/auth"
	"webapp/entity"
)

type DailyRecordService struct {
	repo DailyRecordRepository
}

func NewDailyRecordService(repo DailyRecordRepository) *DailyRecordService {
	return &DailyRecordService{repo: repo}
}

func (us *DailyRecordService) RegisterDailyRecord(ctx context.Context, record *entity.DailyRecord) (*entity.DailyRecord, error) {
	uid, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, errors.New("user id not found")
	}
	record.UserId = uid

	record, err := us.repo.Register(uid, record)
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (us *DailyRecordService) GetDailyRecordList(ctx context.Context, limit, offset int) ([]entity.DailyRecord, error) {
	uid, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, errors.New("user id not found")
	}

	records, err := us.repo.GetList(uid, limit, offset)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (us *DailyRecordService) GetDailyRecordById(ctx context.Context, recordId int64) (*entity.DailyRecord, error) {
	uid, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, errors.New("user id not found")
	}

	record, err := us.repo.GetById(uid, recordId)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (us *DailyRecordService) UpdateDailyRecord(ctx context.Context, record *entity.DailyRecord) (*entity.DailyRecord, error) {
	uid, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, errors.New("user id not found")
	}
	record.UserId = uid

	updatedDailyRecord, err := us.repo.Update(uid, record)
	if err != nil {
		return nil, err
	}
	return updatedDailyRecord, nil
}

func (us *DailyRecordService) DeleteDailyRecord(ctx context.Context, recordId int64) error {
	uid, ok := auth.GetUserID(ctx)
	if !ok {
		return errors.New("user id not found")
	}

	if err := us.repo.Delete(uid, recordId); err != nil {
		return err
	}
	return nil
}
