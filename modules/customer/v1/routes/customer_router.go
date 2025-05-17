package customer_routes

import (
	customer_handlers "divine-pizzeria-backend/modules/customer/v1/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterCustomerRoutes(router *gin.RouterGroup, db *gorm.DB) {
	h := customer_handlers.NewCustomerHandler(db)
	customer := router.Group("/customer/v1")
	{
		customer.GET("/", h.ListCustomers)
		customer.POST("/", h.AddCustomer)
	}
}
