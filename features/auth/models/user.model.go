package models

type UserModel struct {
	ID        int64  `json:"id" db:"id"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	Name      string `json:"name" db:"name"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
