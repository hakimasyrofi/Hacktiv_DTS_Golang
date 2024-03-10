package routers

import (
	"assignment2/handlers"
	"assignment2/repositories"
	"assignment2/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func OrderRoutes(router *gin.Engine, db *gorm.DB) {
	repository := repositories.NewRepository(db)
	service := services.NewService(repository)
	handler := handlers.NewHandler(service)

	router.GET("/orders", handler.GetAllOrders)
	router.POST("/orders", handler.CreateOrder)
	router.PUT("/orders/:orderID", handler.UpdateOrder)
	router.DELETE("/orders/:orderID", handler.DeleteOrder)
}
