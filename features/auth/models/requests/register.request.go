package requests

type RegisterRequest struct {
	Email    string `json:"email" binding:"required" validate:"email"`
	Password string `json:"password" binding:"required" validate:"min=8,max=32"`
	Name     string `json:"name" binding:"required" validate:"min=3,max=32"`
}
