package handlers

import (
	"deps/url-shortener/internal/constants"
	"deps/url-shortener/internal/database"
	"deps/url-shortener/internal/models"
	"deps/url-shortener/internal/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func ShortenendUrlHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	originalUrl, err := database.GetOriginalUrlByCode(id)
	if err == nil {
		ctx.Redirect(http.StatusPermanentRedirect, originalUrl)
		return
	}
	if err == mongo.ErrNoDocuments {
		ctx.IndentedJSON(http.StatusNotFound, constants.UrlNotFoundError)
	} else {
		ctx.IndentedJSON(http.StatusInternalServerError, constants.InternalServerError)
	}
}

func ShortenUrlHandler(ctx *gin.Context) {
	var requestBody models.UrlShortenRequest
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, constants.InvalidRequestBodyError)
		return
	}
	if !utils.IsUrlValid(requestBody.OringinalUrl) {
		ctx.IndentedJSON(http.StatusBadRequest, constants.InvalidUrlFormatError)
		return
	}
	code, err := database.GenerateUrlCode(requestBody.OringinalUrl)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, constants.InternalServerError)
	} else {
		ctx.IndentedJSON(http.StatusCreated, gin.H{"short_url": fmt.Sprintf("%s/%s", constants.BaseUrl, code)})
	}
}
