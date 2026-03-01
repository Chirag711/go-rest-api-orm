package routes

import (
	"github.com/gin-gonic/gin"
	"go-rest-api-orm/controllers"
)

func SetupRoutes(router *gin.Engine) {

	api := router.Group("/api")

	{
		api.POST("/users", controllers.CreateUser)
		api.GET("/users", controllers.GetUsers)
		api.GET("/users/:id", controllers.GetUser)
		api.PUT("/users/:id", controllers.UpdateUser)
		api.DELETE("/users/:id", controllers.DeleteUser)
	}

}