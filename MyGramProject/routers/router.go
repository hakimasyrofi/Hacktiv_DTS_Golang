package routers

import (
	"MyGramProject/handlers"
	"MyGramProject/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler handlers.UserHandler) *gin.Engine {
	router := gin.Default()

	userGroup := router.Group("/users", middleware.AuthMiddleware())
	{
		userGroup.POST("/register", userHandler.RegisterUser)
		userGroup.POST("/login", userHandler.LoginUser)
		userGroup.PUT("/", userHandler.UpdateUser)
		userGroup.DELETE("/", userHandler.DeleteUser)
	}

	return router
}
