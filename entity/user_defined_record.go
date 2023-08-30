package entity

import "time"

type UserDefinedRecord struct {
	Id        int64     `json:"id"`
	UserId    string    `json:"user_id"`
	ItemName  string    `json:"item_name"`
	UnitType  UnitType  `json:"unit_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
