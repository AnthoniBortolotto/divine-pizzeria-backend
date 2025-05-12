package pizza_routes

import (
	pizza_handlers "divine-pizzeria-backend/modules/pizza/v1/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterPizzaRoutes(router *gin.RouterGroup, db *gorm.DB) {

	h := pizza_handlers.NewPizzaHandler(db)

	pizza := router.Group("/pizza/v1")
	{
		pizza.GET("/sizes", h.ListPizzaSizes)
		pizza.POST("/sizes", h.AddPizzaSize)
		pizza.GET("/flavors", h.ListPizzaFlavors)
		pizza.POST("/flavors", h.AddPizzaFlavor)
	}
}
