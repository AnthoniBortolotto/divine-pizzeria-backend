package pizza_routes

import (
	auth_middleware "divine-pizzeria-backend/modules/auth/v1/middleware"
	pizza_handlers "divine-pizzeria-backend/modules/pizza/v1/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterPizzaRoutes(router *gin.RouterGroup, db *gorm.DB) {
	h := pizza_handlers.NewPizzaHandler(db)

	pizza := router.Group("/pizza/v1")
	{
		// Public routes
		pizza.GET("/sizes", h.ListPizzaSizes)
		pizza.GET("/flavors", h.ListPizzaFlavors)

		// Admin-only routes
		admin := pizza.Group("")
		admin.Use(auth_middleware.AuthMiddleware(), auth_middleware.AdminOnly())
		{
			admin.POST("/sizes", h.AddPizzaSize)
			admin.POST("/flavors", h.AddPizzaFlavor)
		}
	}
}
