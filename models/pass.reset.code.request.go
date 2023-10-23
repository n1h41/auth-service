package models

type PassResetCodeRequest struct {
	UserId     int64 `json:"user_id"`
	ResetCode string `json:"reset_code"`
}
