package route

import (
	"github.com/gin-gonic/gin"
	"go-authentication-template/handler"
)

type AuthRouteHandler struct {
	authHandler handler.AuthHandler
}

func NewAuthHandler(authHandler handler.AuthHandler) AuthRouteHandler {
	return AuthRouteHandler{authHandler: authHandler}
}

func (auth *AuthRouteHandler) AuthRoute(group *gin.RouterGroup) {
	router := group.Group("auth")
	router.POST("/register", auth.authHandler.Register)
	router.POST("/login", auth.authHandler.Login)
}
