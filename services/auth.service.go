package services

import (
	"encoding/json"
	"n1h41/auth-service/models"

	"github.com/jmoiron/sqlx"
)

type AuthService interface {
	GetUserDetails(email string) models.UserModel
	RegisterUser(data *models.RegisterRequest) (databaseResponse models.DatabaseResponse)
	CheckIfUserExists(email string) bool
	StorePasswordResetCode(data *models.PasswordResetRequest) models.DatabaseResponse
}

type authService struct {
	db *sqlx.DB
}

func (s *authService) RegisterUser(data *models.RegisterRequest) (databaseResponse models.DatabaseResponse) {
	inputData, err := json.Marshal(data)
	if err != nil {
		println(err)
	}
	inputDataString := string(inputData)
	row := s.db.QueryRowx("select * from auth_service__insert_user($1::jsonb)", inputDataString)
	if err := row.StructScan(&databaseResponse); err != nil {
		println(err.Error())
	}
	return databaseResponse
}

func (s *authService) GetUserDetails(email string) models.UserModel {
	var user models.UserModel
	row := s.db.QueryRowx("SELECT * FROM users WHERE email = $1", email)
	if err := row.StructScan(&user); err != nil {
		println(err.Error())
	}
	return user
}

func (s *authService) CheckIfUserExists(email string) bool {
	var count int
	err := s.db.Get(&count, "SELECT COUNT(*) FROM users WHERE email = $1", email)
	if err != nil {
		println(err.Error())
		return false
	}
	if count > 0 && count < 2 {
		return true
	}
	return false
}

func (s *authService) StorePasswordResetCode(data *models.PasswordResetRequest) (databaseResponse models.DatabaseResponse) {
	inputData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	row := s.db.QueryRowx("select * from auth_service__reset_pass($1::jsonb)", string(inputData))
	if err := row.StructScan(&databaseResponse); err != nil {
		println(err.Error())
	}
	return databaseResponse
}

// INFO: NewAuthService returns a new instance of AuthService
func NewAuthService(db *sqlx.DB) AuthService {
	var service AuthService
	service = &authService{
		db: db,
	}
	return service
}
