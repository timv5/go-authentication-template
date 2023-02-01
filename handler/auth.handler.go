package handler

import (
	"github.com/gin-gonic/gin"
	"go-authentication-template/configs"
	"go-authentication-template/dto/request"
	"go-authentication-template/dto/response"
	"go-authentication-template/models"
	"go-authentication-template/repository"
	"go-authentication-template/service"
	"go-authentication-template/utils"
	"gorm.io/gorm"
	"net/http"
)

type AuthHandler struct {
	postgresDB              *gorm.DB
	authService             *service.AuthService
	userRepository          *repository.UserRepository
	emailTemplateRepository *repository.EmailTemplateRepository
	userEmailRepository     *repository.UserEmailTemplateRepository
	config                  *configs.Config
}

func NewAuthHandler(postgresDB *gorm.DB, authService *service.AuthService, userRepository *repository.UserRepository,
	emailTemplateRepository *repository.EmailTemplateRepository, userEmailRepository *repository.UserEmailTemplateRepository, config *configs.Config) AuthHandler {
	return AuthHandler{
		postgresDB:              postgresDB,
		authService:             authService,
		userRepository:          userRepository,
		emailTemplateRepository: emailTemplateRepository,
		userEmailRepository:     userEmailRepository,
		config:                  config,
	}
}

func (authHandler AuthHandler) Register(ctx *gin.Context) {
	var registerPayload *request.Register
	if err := ctx.ShouldBindJSON(&registerPayload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	var user models.User

	// check if user exists with username
	user = authHandler.userRepository.GetUserByUsername(registerPayload.Username)
	if (models.User{} != user) {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Username already exists"})
		return
	}

	// check if user exists with email
	user = authHandler.userRepository.GetUserByEmail(registerPayload.Email)
	if (models.User{} != user) {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Email already exists"})
		return
	}

	// check if passwords matches
	if registerPayload.Password != registerPayload.PasswordConfirm {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "Passwords don't match"})
		return
	}

	hashedPassword, err := utils.HashPassword(registerPayload.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	var savedUser models.User
	savedUser, err = authHandler.userRepository.SaveUser(registerPayload.Email, hashedPassword, registerPayload.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Registration error, please contact support"})
		return
	}

	userResponse := &response.UserResponse{
		ID:       savedUser.ID,
		Email:    savedUser.Email,
		Username: savedUser.Username,
	}

	// try to send email
	utils.SendEmail(&savedUser, authHandler.emailTemplateRepository, authHandler.userEmailRepository, authHandler.config)

	// success
	ctx.JSON(http.StatusCreated, gin.H{"user": userResponse})
}

func (authHandler AuthHandler) Login(ctx *gin.Context) {
	var loginPayload *request.Login
	if err := ctx.ShouldBindJSON(&loginPayload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Missing request params"})
		return
	}

	var user models.User
	user = authHandler.userRepository.GetUserByEmail(loginPayload.Email)
	if (models.User{} == user) {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "User does not exist"})
		return
	}

	if err := utils.VerifyPassword(user.Password, loginPayload.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Incorrect password"})
		return
	}

	accessToken, err := authHandler.authService.GenerateToken(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "success", "message": "Logged in"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "payload": accessToken})
	}
}
