package handlers

import (
	"assignment2/entities"
	"assignment2/schemas"
	"assignment2/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service services.Service
}

func NewHandler(service services.Service) *handler {
	return &handler{service}
}

func (h *handler) GetAllOrders(c *gin.Context) {
	items, err := h.service.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Error")
	}
	c.JSON(http.StatusOK, items)
}

func (h *handler) CreateOrder(c *gin.Context){
	var orderInput schemas.OrderInput

	if err := c.BindJSON(&orderInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newOrder := entities.Order{
		OrderID:      0,
		CustomerName: orderInput.CustomerName,
		OrderedAt:    orderInput.OrderedAt,
	}

	err := h.service.CreateOrder(newOrder, orderInput.Items)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully"})
}

func (h *handler) UpdateOrder(c *gin.Context) {
	var orderInput schemas.OrderInputUpdate
	if err := c.ShouldBindJSON(&orderInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	orderID, err := strconv.Atoi(c.Param("orderID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	err = h.service.UpdateOrder(orderID, orderInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully"})
}

func (h *handler) DeleteOrder(c *gin.Context) {
	orderID := c.Param("orderID")
	orderIDInt, err := strconv.Atoi(orderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	err = h.service.DeleteOrderWithItems(orderIDInt)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Error deleting order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
