package dailyRecord

import "webapp/entity"

type RegisterDailyRecordRequest struct {
	Weather         string                  `json:"weather" validate:"required"`
	Memo            string                  `json:"memo" validate:"required"`
	OptionalRecords []OptionalRecordRequest `json:"optional_records"`
}

type GetDailyRecordByIdRequest struct {
	Id int64 `json:"id" validate:"required"`
}

type UpdateDailyRecordRequest struct {
	Id              int64                   `json:"id" validate:"required"`
	Weather         string                  `json:"weather" validate:"required"`
	Memo            string                  `json:"memo" validate:"required"`
	OptionalRecords []OptionalRecordRequest `json:"optional_records"`
}

type DeleteDailyRecordRequest struct {
	Id int64 `json:"id" validate:"required"`
}

type OptionalRecordRequest struct {
	Id            int64   `json:"id,omitempty" validate:"required"`
	UserDefinedId int64   `json:"user_defined_id" validate:"required"`
	Value         float64 `json:"value" validate:"required"`
}

type DailyRecordResponse struct {
	DailyRecord     entity.DailyRecord      `json:"daily_record"`
	OptionalRecords []entity.OptionalRecord `json:"optional_records"`
}
