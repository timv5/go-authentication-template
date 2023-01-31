package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-authentication-template/configs"
	"go-authentication-template/handler"
	"go-authentication-template/route"
	"go-authentication-template/service"
	"log"
	"os"
)

var (
	server              *gin.Engine
	UserController      handler.UserHandler
	UserRouteController route.UserRouteHandler
	AuthController      handler.AuthHandler
	AuthRouteController route.AuthRouteHandler
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic("Could not initialize app")
	}

	err = os.Setenv("JWT_SECRET", config.JwtSecret)
	if err != nil {
		panic("Could not initialize app")
		return
	}

	// connect to database
	configs.ConnectToDB(&config)

	// initialize services
	authService := service.NewAuthService(&config)

	// initialize controllers and routes
	UserController = handler.NewUserHandler(configs.DB, authService)
	UserRouteController = route.NewUserRouteHandler(UserController)
	AuthController = handler.NewAuthHandler(configs.DB, authService)
	AuthRouteController = route.NewAuthHandler(AuthController)

	server = gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8080", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	UserRouteController.UserRoute(router)
	AuthRouteController.AuthRoute(router)

	log.Fatal(server.Run(":" + config.ServerPort))
}
