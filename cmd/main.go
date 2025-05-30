package main

import (
	"go-api/internal/controllers"
	"go-api/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	server := gin.Default()

	userController := controllers.NewUserController(database.DB)
	server.GET("/usuario/:id", userController.GetUser)
	server.GET("/usuarios", userController.GetAllUsers)
	server.POST("/usuario", userController.CreateUser)
	server.Run(":3334")
}
