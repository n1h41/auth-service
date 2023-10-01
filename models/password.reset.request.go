package models

type PasswordResetRequest struct {
	Email     string `json:"email"`
	ResetCode string `json:"reset_code"`
}
