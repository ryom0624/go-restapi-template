package service

import (
	"context"
	"errors"
	"webapp/auth"
	"webapp/entity"
)

type UserDefinedRecordService struct {
	repo UserDefinedRecordRepository
}

func NewUserDefinedRecordService(repo UserDefinedRecordRepository) *UserDefinedRecordService {
	return &UserDefinedRecordService{repo: repo}
}

func (udrs *UserDefinedRecordService) RegisterUserDefinedRecord(ctx context.Context, record *entity.UserDefinedRecord) (*entity.UserDefinedRecord, error) {
	uid, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, errors.New("user id not found")
	}
	record.UserId = uid

	record, err := udrs.repo.Register(uid, record)
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (udrs *UserDefinedRecordService) GetUserDefinedRecordList(ctx context.Context, limit, offset int) ([]entity.UserDefinedRecord, error) {
	uid, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, errors.New("user id not found")
	}

	records, err := udrs.repo.GetList(uid, limit, offset)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (udrs *UserDefinedRecordService) GetUserDefinedRecordById(ctx context.Context, recordId int64) (*entity.UserDefinedRecord, error) {
	uid, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, errors.New("user id not found")
	}

	user, err := udrs.repo.GetById(uid, recordId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (udrs *UserDefinedRecordService) UpdateUserDefinedRecord(ctx context.Context, record *entity.UserDefinedRecord) (*entity.UserDefinedRecord, error) {
	uid, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, errors.New("user id not found")
	}
	record.UserId = uid

	updatedUserDefinedRecord, err := udrs.repo.Update(uid, record)
	if err != nil {
		return nil, err
	}
	return updatedUserDefinedRecord, nil
}

func (udrs *UserDefinedRecordService) DeleteUserDefinedRecord(ctx context.Context, id int64) error {
	uid, ok := auth.GetUserID(ctx)
	if !ok {
		return errors.New("user id not found")
	}

	if err := udrs.repo.Delete(uid, id); err != nil {
		return err
	}
	return nil
}
