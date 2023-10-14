package models

type PassResetCodeRequest struct {
	Email     string `json:"email"`
	ResetCode string `json:"reset_code"`
}
