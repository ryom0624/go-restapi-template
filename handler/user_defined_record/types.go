package userDefinedRecord

import (
	"webapp/entity"
)

type RegisterUserDefinedRecordRequest struct {
	ItemName string          `json:"item_name" validate:"required"`
	UnitType entity.UnitType `json:"unit_type" validate:"required,validate_user_defined_record_type"`
}

type GetUserDefinedRecordRequest struct {
	Id int64 `json:"id" validate:"required"`
}

type DeleteUserDefinedRecordRequest struct {
	Id int64 `json:"id" validate:"required"`
}
