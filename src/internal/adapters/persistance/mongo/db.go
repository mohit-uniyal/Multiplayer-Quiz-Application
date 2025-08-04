package mongo_persistance

import (
	"context"
	"fmt"
	"log"
	"multiplayer-quiz-application/src/internal/config"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Database struct {
	DB *mongo.Database
}

func NewDatabase(config *config.Config) (*Database, error) {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(config.MONGO_CONNECTION_STRING).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongo database: %w", err)
	}

	database := client.Database(config.MONGO_DATABASE_NAME)

	log.Printf("successfully connected to mongo db")

	return &Database{DB: database}, nil

}

func (db *Database) Close() error {
	return db.DB.Client().Disconnect(context.Background())
}

func (db *Database) GetDB() *mongo.Database {
	return db.DB
}
