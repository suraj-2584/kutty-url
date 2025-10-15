package main

import (
	"deps/url-shortener/internal/handlers"
	"deps/url-shortener/internal/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadSecrets()
	router := gin.Default()
	router.Use()
	router.GET("/:id", handlers.ShortenendUrlHandler)
	router.POST("/shorten", handlers.ShortenUrlHandler)
	router.GET("/health", handlers.HealthHandler)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:*", "https://kutty-url.in"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowWildcard:    true,
	}))
	router.Run(":5001")
}
