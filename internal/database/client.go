package database

import (
	"context"
	"deps/url-shortener/internal/models"
	"deps/url-shortener/internal/utils"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// check code for GET calls
func GetOriginalUrlByCode(code string) (string, error) {
	urlsDocument, err := fetchUrlsDocument(bson.M{"code": code})
	return urlsDocument.OringinalUrl, err
}

func GenerateUrlCode(url string) (string, error) {
	if code := getCodeByOriginalUrl(url); code != "" {
		return code, nil
	}
	mongoDb := InitClient()
	var sequenceDocument models.SequenceDocument
	var code string
	err := mongoDb.sequenceCollection.FindOneAndUpdate(
		context.Background(),
		bson.M{"_id": "short_url"},
		bson.M{"$inc": bson.M{"counter": 1}},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&sequenceDocument)
	if err == nil {
		code = utils.IntToBase62String(sequenceDocument.Counter)
		insertDocument := models.UrlsDocument{Code: code, OringinalUrl: url, CreatedAt: time.Now()}
		_, insertError := mongoDb.urlsCollection.InsertOne(context.Background(), insertDocument)
		if insertError != nil {
			err = insertError
		}
	}
	return code, err
}

func getCodeByOriginalUrl(originalUrl string) string {
	urlsDocument, _ := fetchUrlsDocument(bson.M{"original_url": originalUrl})
	return urlsDocument.Code
}
func fetchUrlsDocument(filter map[string]any) (models.UrlsDocument, error) {
	mongoDb := InitClient()
	var urlDocument models.UrlsDocument
	err := mongoDb.urlsCollection.FindOne(context.Background(), filter).Decode(&urlDocument)
	return urlDocument, err
}
