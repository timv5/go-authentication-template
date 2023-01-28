package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-authentication-template/configs"
	"go-authentication-template/handler"
	"go-authentication-template/route"
	"log"
	"net/http"
)

var (
	server              *gin.Engine
	UserController      handler.UserHandler
	UserRouteController route.UserRouteHandler
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic("Could not initialize app")
	}

	// connect to database
	configs.ConnectToDB(&config)

	// initialize controllers
	UserController = handler.NewUserHandler(configs.DB)
	UserRouteController = route.NewUserRouteHandler(UserController)

	server = gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8080", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	UserRouteController.UserRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}
