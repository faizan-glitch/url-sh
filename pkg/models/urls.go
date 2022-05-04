package models

import (
	"context"

	"github.com/faizan-glitch/url-sh/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
)

type UrlSchema struct {
	LongUrl   string `json:"long_url" bson:"long_url"`
	ShortUrl  string `json:"short_url" bson:"short_url"`
	ShortCode string `json:"short_code" bson:"short_code"`
	Password  string `json:"password" bson:"password"`
}

func (u *UrlSchema) Insert() error {
	_, err := database.Db.Collection("Urls").InsertOne(context.Background(), u)

	if err != nil {
		return err
	}

	return nil
}

func (u *UrlSchema) FindByShortCode(shortCode string) error {

	if err := database.Db.Collection("Urls").FindOne(context.TODO(), bson.M{"short_code": shortCode}).Decode(&u); err != nil {
		return err
	}

	return nil
}
