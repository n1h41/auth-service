package requests

type ResetPassRequest struct {
	Email string `json:"email" binding:"required" validate:"required,email"`
}
