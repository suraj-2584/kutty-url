package main

import (
	"deps/url-shortener/internal/handlers"
	"deps/url-shortener/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadSecrets()
	router := gin.Default()
	router.GET("/:id", handlers.ShortenendUrlHandler)
	router.POST("/shorten", handlers.ShortenUrlHandler)
	router.Run("localhost:5001")
}
