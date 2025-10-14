package models

import (
	"time"
)

// API request model
type UrlShortenRequest struct {
	OringinalUrl string `json:"original_url"`
}

// MongoDB Document models
type UrlsDocument struct {
	OringinalUrl string    `bson:"original_url"`
	Code         string    `bson:"code"`
	CreatedAt    time.Time `bson:"created_at"`
}

type SequenceDocument struct {
	Id      string `bson:"_id"`
	Counter int32  `bson:"counter"`
}
