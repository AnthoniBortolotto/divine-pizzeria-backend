package order_routes

import (
	auth_middleware "divine-pizzeria-backend/modules/auth/v1/middleware"
	order_handlers "divine-pizzeria-backend/modules/order/v1/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterOrderRoutes(router *gin.RouterGroup, db *gorm.DB) {
	h := order_handlers.NewOrderHandler(db)

	order := router.Group("/order/v1")
	{
		// Protected routes
		order.Use(auth_middleware.AuthMiddleware())
		{
			order.GET("/", h.ListOrders)
			order.POST("/", h.CreateOrder)
		}
	}
}
