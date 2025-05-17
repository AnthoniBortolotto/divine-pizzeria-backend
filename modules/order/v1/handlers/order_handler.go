package order_handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderHandler struct{}

func NewOrderHandler(db *gorm.DB) *OrderHandler {
	return &OrderHandler{}
}

func (h *OrderHandler) ListOrders(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "List of orders",
	})
}

func (h *OrderHandler) AddOrder(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "Order added successfully",
	})
}
