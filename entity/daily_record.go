package entity

import "time"

type DailyRecord struct {
	Id              int64            `json:"id"`
	UserId          string           `json:"user_id"`
	Weather         string           `json:"weather"`
	Memo            string           `json:"memo"`
	OptionalRecords []OptionalRecord `json:"optional_records" gorm:"hasMany:OptionalRecord;foreignKey:DailyRecordId"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
}

type OptionalRecord struct {
	Id            int64     `json:"id"`
	DailyRecordId int64     `json:"daily_record_id"`
	UserDefinedId int64     `json:"user_defined_id"`
	Value         float64   `json:"value"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
