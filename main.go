package main

import (
	"deps/url-shortener/internal/handlers"
	"deps/url-shortener/internal/utils"
	"strings"

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
		AllowOriginFunc: func(origin string) bool {
			return strings.HasPrefix(origin, "http://localhost")
		},
		AllowCredentials: true,
	}))
	router.Run(":5001")
}
