package services

import (
	"encoding/json"
	"fmt"
	"n1h41/auth-service/models"

	"github.com/jmoiron/sqlx"
)

type AuthService interface {
	GetUserDetails(email string) (user *models.UserModel, err error)
	RegisterUser(data *models.RegisterRequest) (databaseResponse models.DatabaseResponse)
	CheckIfUserExists(email string) bool
	StorePasswordResetCode(reset_code string, user_id int) models.DatabaseResponse
	ResetPassword(data *models.PassResetCodeRequest) models.DatabaseResponse
}

type authService struct {
	db *sqlx.DB
}

func (s *authService) RegisterUser(data *models.RegisterRequest) (databaseResponse models.DatabaseResponse) {
	inputData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
	}
	inputDataString := string(inputData)
	row := s.db.QueryRowx("SELECT * FROM auth_service_register_user($1::jsonb)", inputDataString)
	if err := row.StructScan(&databaseResponse); err != nil {
		fmt.Println(err.Error())
	}
	return databaseResponse
}

func (s *authService) GetUserDetails(email string) (user *models.UserModel, err error) {
	user = &models.UserModel{}
	row := s.db.QueryRowx("SELECT * FROM users WHERE email = $1", email)
	if err := row.StructScan(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *authService) CheckIfUserExists(email string) bool {
	var count int
	err := s.db.Get(&count, "SELECT 1 FROM users WHERE email = $1", email)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if count > 0 && count < 2 {
		return true
	}
	return false
}

// Store password reset code
func (s *authService) StorePasswordResetCode(reset_code string, user_id int) (databaseResponse models.DatabaseResponse) {
	_, err := s.db.Exec("INSERT INTO reset_pass (reset_code, user_id) VALUES ($1, $2)", reset_code, user_id)
	if err != nil {
		databaseResponse.Status = false
		databaseResponse.Message = err.Error()
	}
	databaseResponse.Status = true
	databaseResponse.Message = "Password reset code stored"
	return databaseResponse
}

func (s *authService) ResetPassword(data *models.PassResetCodeRequest) (databaseResponse models.DatabaseResponse) {
	databaseRequestParams, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	databaseRequestParamsString := string(databaseRequestParams)
	row := s.db.QueryRowx("SELECT * FROM auth_service_reset_pass($1::jsonb)", databaseRequestParamsString)
	if err := row.StructScan(&databaseResponse); err != nil {
		fmt.Println(err.Error())
	}
	return databaseResponse
}

// Creates and returns a new instance of AuthService
func NewAuthService(db *sqlx.DB) AuthService {
	var service AuthService
	service = &authService{
		db: db,
	}
	return service
}
