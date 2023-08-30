package user

type RegisterUserRequest struct {
	Birthday      int64  `json:"birthday" validate:"required"`
	LastName      string `json:"last_name" validate:"required"`
	FirstName     string `json:"first_name" validate:"required"`
	LastNameKana  string `json:"last_name_kana" validate:"required"`
	FirstNameKana string `json:"first_name_kana" validate:"required"`
	Sex           int    `json:"sex" validate:"required"`
	Prefecture    string `json:"prefecture" validate:"required"`
}

type UpdateUserRequest struct {
	Birthday      int64  `json:"birthday" validate:"required"`
	LastName      string `json:"last_name" validate:"required"`
	FirstName     string `json:"first_name" validate:"required"`
	LastNameKana  string `json:"last_name_kana" validate:"required"`
	FirstNameKana string `json:"first_name_kana" validate:"required"`
	Sex           int    `json:"sex" validate:"required"`
	Prefecture    string `json:"prefecture" validate:"required"`
}
