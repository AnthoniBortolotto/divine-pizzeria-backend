package auth_routes

import (
	auth_handlers "divine-pizzeria-backend/modules/auth/v1/handlers"
	auth_middleware "divine-pizzeria-backend/modules/auth/v1/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAuthRoutes(router *gin.Engine, db *gorm.DB) {
	authHandler := auth_handlers.NewAuthHandler(db)

	auth := router.Group("/api/v1/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)

		// Protected routes
		protected := auth.Group("")
		protected.Use(auth_middleware.AuthMiddleware())
		{
			// Add protected routes here
		}

		// Admin only routes
		admin := auth.Group("/admin")
		admin.Use(auth_middleware.AuthMiddleware(), auth_middleware.AdminOnly())
		{
			// Add admin only routes here
		}
	}
}
