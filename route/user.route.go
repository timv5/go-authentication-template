package route

import (
	"github.com/gin-gonic/gin"
	"go-authentication-template/handler"
	"go-authentication-template/middleware"
)

type UserRouteHandler struct {
	userHandler handler.UserHandler
}

func NewUserRouteHandler(userHandler handler.UserHandler) UserRouteHandler {
	return UserRouteHandler{userHandler: userHandler}
}

func (ur *UserRouteHandler) UserRoute(group *gin.RouterGroup) {
	router := group.Group("users")
	router.Use(middleware.Authenticate())
	router.POST("/getByEmail", ur.userHandler.GetByEmail)
	router.GET("/getByUsername", ur.userHandler.GetByUsername)
}
