package models

type PassResetCodeRequest struct {
	UserId      int64  `json:"user_id" validate:"required"`
	ResetCode   string `json:"reset_code" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}
