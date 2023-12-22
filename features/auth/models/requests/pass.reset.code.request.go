package requests

type PassResetCodeRequest struct {
	UserId      int64  `json:"userId" validate:"required"`
	ResetCode   string `json:"resetCode" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
}
