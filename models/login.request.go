package models

type LoginRequest struct {
  Email    string `json:"email" binding:"required" validata:"email"`
  Password string `json:"password" binding:"required" validate:"min=8,max=32"`
}
