package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/ridhohrnf/task-5-vix-btpns-Muhammad-Ridho-Heranof/controllers"
	"github.com/ridhohrnf/task-5-vix-btpns-Muhammad-Ridho-Heranof/middlewares"
)

func SetupUserRoutes(router *gin.Engine, db *gorm.DB) {
	userController := controllers.NewUserController(db)

	router.POST("/users/register", userController.Register)
	router.POST("/users/login", userController.Login)

	authGroup := router.Group("/users")
	authGroup.Use(middlewares.AuthMiddleware())
	{
		authGroup.PUT("/:userId", userController.UpdateUser)
		authGroup.DELETE("/:userId", userController.DeleteUser)
	}
}
