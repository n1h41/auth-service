package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gopkg.in/mail.v2"

	"n1h41/auth-service/features/auth/models/requests"
	"n1h41/auth-service/features/auth/services"
	"n1h41/auth-service/utils"
)

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	GetResetPasswordLink(c *gin.Context)
	ResetPassword(c *gin.Context)
	Status(c *gin.Context)
}

type authController struct {
	authService services.AuthService
}

func (controller *authController) Register(c *gin.Context) {
	var data requests.RegisterRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data.Password = string(utils.HashString(data.Password))
	userExists := controller.authService.CheckIfUserExists(data.Email)
	if userExists {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": false, "message": "User already exists"})
		return
	}
	serviceResponse := controller.authService.RegisterUser(&data)
	if serviceResponse.Status == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": false, "message": serviceResponse.Message})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"status": true, "message": "User created", "data": data})
	return
}

func (controller *authController) Login(c *gin.Context) {
	var data requests.LoginRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userExists := controller.authService.CheckIfUserExists(data.Email)
	if !userExists {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": false, "message": "User doesn't exist! Register"})
		return
	}
	user, _ := controller.authService.GetUserDetails(data.Email)
	if validPassword := utils.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); !validPassword {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid password"})
		return
	}
	token, err := utils.CreateJwtToken(*user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": false, "message": "Error creating jwt token"})
		return
	}
	//
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Login successful", "status": true, "data": gin.H{"token": token}})
}

func (controller *authController) GetResetPasswordLink(c *gin.Context) {
	var data requests.ResetPassRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := controller.authService.GetUserDetails(data.Email)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": false, "message": "User doesn't exist! Register"})
		return
	}

	uuid := utils.GenerateUUID()
	resetLink := "http://localhost:8080/reset-password/" + uuid + "?userId=" + strconv.Itoa(int(user.ID))

	databaseRequest := requests.PassResetCodeRequest{
		UserId:    user.ID,
		ResetCode: uuid,
	}

	// INFO: Saving pass reset code to database
	databaseResponse := controller.authService.StorePasswordResetCode(databaseRequest.ResetCode, int(user.ID))

	if databaseResponse.Status == false {
		log.Panicln(databaseResponse.Message)
	}

	// INFO: Sending mail
	mail := mail.NewMessage()
	mail.SetHeader("From", "auth-service@gmail.com")
	mail.SetHeader("To", user.Email)
	mail.SetHeader("Subject", "Reset password link")
	mail.SetBody("text/html", "<h1>Reset password link</h1><p>Click on the link to reset your password</p><a href='"+resetLink+"'>Reset password</a>")

	if err := utils.SendMail(mail); err != nil {
		log.Panicln(err)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"status": true, "message": "Reset password link sent to your email"})
}

func (controller *authController) ResetPassword(c *gin.Context) {
	// NOTE: reset password logic
	userId, err := strconv.ParseInt(c.Query("userId"), 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid user id"})
		return
	}
	databaseRequest := requests.PassResetCodeRequest{
		UserId:      userId,
		ResetCode:   c.Param("resetCode"),
		NewPassword: c.Query("newPassword"),
	}
	databaseResponse := controller.authService.ResetPassword(&databaseRequest)
	if databaseResponse.Status == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": false, "message": databaseResponse.Message})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"status": true, "message": "Password reset successfull"})
}

func (controller *authController) Status(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Auth endpoint"})
	return
}

// INFO: NewAuthController is a constructor for authController
func NewAuthController(router *gin.Engine, authService services.AuthService) {
	controller := &authController{
		authService: authService,
	}
	api := router.Group("auth")
	{
		api.GET("/status", controller.Status)
		api.POST("/register", controller.Register)
		api.POST("/login", controller.Login)
		api.POST(("/reset-password"), controller.GetResetPasswordLink)
		api.GET(("/reset-password/:resetCode"), controller.ResetPassword)
	}
}
