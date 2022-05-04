package database

import (
	"context"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Database

func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	credential := options.Credential{
		AuthSource: "admin",
		Username:   viper.GetString("DB_USER"),
		Password:   viper.GetString("DB_PASS"),
	}

	clientOpts := options.Client().ApplyURI(viper.GetString("DB_URL")).
		SetAuth(credential)

	client, conErr := mongo.Connect(ctx, clientOpts)

	if conErr != nil {
		panic(conErr)
	}

	Db = client.Database(viper.GetString("DB_NAME"))

	Db.Collection("Urls").Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys:    bson.D{{Key: "short_code", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
}
