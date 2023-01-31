package handler

import (
	"github.com/gin-gonic/gin"
	"go-authentication-template/dto/request"
	"go-authentication-template/models"
	"go-authentication-template/service"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type UserHandler struct {
	postgresDB  *gorm.DB
	authService *service.AuthService
}

func NewUserHandler(postgresDB *gorm.DB, authService *service.AuthService) UserHandler {
	return UserHandler{
		postgresDB:  postgresDB,
		authService: authService,
	}
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
