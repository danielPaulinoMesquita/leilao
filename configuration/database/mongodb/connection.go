package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"leilao/configuration/logger"
	"os"
)

const (
	MONGODB_URL = "MONGODB_URL"
	MONGODB_DB  = "MONGODB_DB"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongoURL := os.Getenv(MONGODB_URL)
	mongoDatabase := os.Getenv(MONGODB_DB)

	logger.Info("MongoDB URL: " + mongoURL)
	logger.Info("MongoDB Database: " + mongoDatabase)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))

	if err != nil {
		logger.Error("Error trying to connect to MongoDB (BANCO DE DADOS CONEXÃO)", err)
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		logger.Error("Error trying to ping MongoDB ", err)
		return nil, err
	}

	return client.Database(mongoDatabase), nil
}
