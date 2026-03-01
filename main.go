package main

import (
	"go-rest-api-orm/config"
	"go-rest-api-orm/models"
	"go-rest-api-orm/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	config.ConnectDB()

	config.DB.AutoMigrate(&models.User{})

	routes.SetupRoutes(router)

	router.Run(":8080")
}