package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/labadiejoaco/home-cloud/backend/handlers"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	r.Use(cors.New(config))

	r.GET("/api/*path", handlers.List)
	r.POST("/api/upload/*path", handlers.Upload)
	r.POST("/api/directory/*path", handlers.CreateDirectory)

	return r
}
