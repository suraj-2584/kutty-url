package database

import (
	"deps/url-shortener/internal/constants"
	"sync"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoDb struct {
	mongoClient        mongo.Client
	sequenceCollection *mongo.Collection
	urlsCollection     *mongo.Collection
}

var (
	m    *MongoDb
	once sync.Once
)

func InitClient() *MongoDb {
	once.Do(func() {
		m = new(MongoDb)
		uri := constants.ConnectionString
		client, err := mongo.Connect(options.Client().
			ApplyURI(uri))
		if err != nil {
			panic(err)
		}
		m.mongoClient = *client
		m.sequenceCollection = client.Database("kutty_url").Collection("sequence")
		m.urlsCollection = client.Database("kutty_url").Collection("urls")
	})
	return m
}
