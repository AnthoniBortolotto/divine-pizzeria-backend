package router

import (
	"divine-pizzeria-backend/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) {
	config.LoadEnv()

	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	RegisterRoutes(router, db)

	port := config.GetEnv("PORT")

	if port == "" {
		panic("PORT environment variable not set")
	}

	router.Run("0.0.0.0:" + port)
}
