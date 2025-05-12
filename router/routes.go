package router

import (
	pizza_routes "divine-pizzeria-backend/modules/pizza/v1/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(
	router *gin.Engine,
	db *gorm.DB,
) {
	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	api := router.Group("/api")
	{
		pizza_routes.RegisterPizzaRoutes(api, db)
	}
}
