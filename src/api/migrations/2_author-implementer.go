package migrations

import (
	"context"
	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	migrate.Register(func(db *mongo.Database) error {

		opt := options.Index().SetName("author")
		keys := bson.D{{"", 1}}
		model := mongo.IndexModel{Keys: keys, Options: opt}
		_, err := db.Collection("my-coll").Indexes().CreateOne(context.TODO(), model)
		if err != nil {
			return err
		}

		return nil
	}, func(db *mongo.Database) error {
		_, err := db.Collection("my-coll").Indexes().DropOne(context.TODO(), "my-index")
		if err != nil {
			return err
		}
		return nil
	})
}
