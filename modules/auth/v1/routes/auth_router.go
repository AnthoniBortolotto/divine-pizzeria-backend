package auth_routes

import (
	auth_handlers "divine-pizzeria-backend/modules/auth/v1/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(router *gin.RouterGroup, db *gorm.DB) {
	h := auth_handlers.NewAuthHandler(db)
	auth := router.Group("/auth/v1")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
	}
}
