package pizza_routes

import (
	auth_middleware "divine-pizzeria-backend/modules/auth/v1/middleware"
	pizza_handlers "divine-pizzeria-backend/modules/pizza/v1/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupPizzaRoutes(router *gin.Engine, db *gorm.DB) {
	pizzaHandler := pizza_handlers.NewPizzaHandler(db)

	pizza := router.Group("/api/v1/pizzas")
	{
		// Public routes
		pizza.GET("/sizes", pizzaHandler.ListPizzaSizes)
		pizza.GET("/flavors", pizzaHandler.ListPizzaFlavors)

		// Admin only routes
		admin := pizza.Group("")
		admin.Use(auth_middleware.AuthMiddleware(), auth_middleware.AdminOnly())
		{
			admin.POST("/sizes", pizzaHandler.AddPizzaSize)
			admin.POST("/flavors", pizzaHandler.AddPizzaFlavor)
		}
	}
}
