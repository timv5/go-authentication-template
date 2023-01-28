package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"go-authentication-template/dto/request"
	"go-authentication-template/dto/response"
	"go-authentication-template/models"
	"go-authentication-template/utils"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
)

type UserHandler struct {
	postgresDB *gorm.DB
}

func NewUserHandler(postgresDB *gorm.DB) UserHandler {
	return UserHandler{postgresDB: postgresDB}
}

func (c UserHandler) GetByEmail(ctx *gin.Context) {
	var userDataPayload *request.UserDataEmail
	if err := ctx.ShouldBindJSON(&userDataPayload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	var user models.User
	resultUser := c.postgresDB.First(&user, "email = ?", strings.ToLower(userDataPayload.Email))
	if resultUser.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid email"})
	} else {
		ctx.JSON(http.StatusOK, user)
	}
}

func (c UserHandler) GetByUsername(ctx *gin.Context) {
	username := ctx.Query("username")
	if username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Missing request param username"})
		return
	}

	var user models.User
	c.postgresDB.Raw("select id, email, password, username, created_at, updated_at from users where username = ?", username).Scan(&user)
	if (models.User{} == user) {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid username"})
	} else {
		ctx.JSON(http.StatusOK, user)
	}
}

func (c UserHandler) Register(ctx *gin.Context) {
	var registerPayload *request.Register
	if err := ctx.ShouldBindJSON(&registerPayload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	var user models.User
	// check if user exists with username
	c.postgresDB.Raw("select id, email, password, username, created_at, updated_at from users where username = ?", registerPayload.Username).Scan(&user)
	if (models.User{} != user) {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Username already exists"})
		return
	}

	// check if user exists with email
	c.postgresDB.Raw("select id, email, password, username, created_at, updated_at from users where email = ?", registerPayload.Email).Scan(&user)
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

	nowTime := time.Now()
	createUser := models.User{
		ID:        uuid.NewV4().String(),
		Email:     registerPayload.Email,
		Password:  hashedPassword,
		Username:  registerPayload.Username,
		CreatedAt: nowTime,
		UpdatedAt: nowTime,
	}

	// save
	savedUser := c.postgresDB.Create(&createUser)
	if savedUser.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Registration error, please contact support"})
		return
	}

	userResponse := &response.UserResponse{
		ID:       createUser.ID,
		Email:    createUser.Email,
		Username: createUser.Username,
	}

	ctx.JSON(http.StatusCreated, gin.H{"user": userResponse})
}

func (c UserHandler) Login(ctx *gin.Context) {
	var loginPayload *request.Login
	if err := ctx.ShouldBindJSON(&loginPayload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Missing request params"})
		return
	}

	var user models.User
	c.postgresDB.Raw("select id, email, password, username, created_at, updated_at from users where email = ?", loginPayload.Email).Scan(&user)
	if (models.User{} == user) {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "User does not exist"})
		return
	}

	if err := utils.VerifyPassword(user.Password, loginPayload.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Incorrect password"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Logged in"})
}
