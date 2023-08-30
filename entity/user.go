package entity

type User struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Birthday      int64  `json:"birthday"`
	LastName      string `json:"last_name"`
	FirstName     string `json:"first_name"`
	LastNameKana  string `json:"last_name_kana"`
	FirstNameKana string `json:"first_name_kana"`
	Sex           int    `json:"sex"`
	Prefecture    string `json:"prefecture"`
	CreatedAt     int64  `json:"created_at"`
	UpdatedAt     int64  `json:"updated_at"`
}
