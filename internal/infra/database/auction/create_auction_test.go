package auction_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"leilao/internal/entity/auction_entity"
	"leilao/internal/infra/database/auction"
	"leilao/internal/internal_error"
	"testing"
	"time"
)

func TestAuctionRepository_CheckAuctionIsClosed(t *testing.T) {
	clientOptions := options.Client().ApplyURI("mongodb://root:root@localhost:27017/auctions?authSource=admin")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	database := client.Database("test_db")
	collection := database.Collection("auctions")

	err = collection.Drop(context.Background())
	if err != nil {
		t.Fatalf("Failed to drop collection: %v", err)
	}

	ar := auction.NewAuctionRepository(database)

	filter := bson.M{"_id": "123"}
	ctx := context.Background()
	duration, err := time.ParseDuration("20s")
	if err != nil {
		t.Fatalf("Failed to parse duration: %v", err)
	}

	t.Run("Expired Auction", func(t *testing.T) {
		expiredAuction := auction.AuctionEntityMongo{
			Id:          "123",
			ProductName: "Test Product",
			Category:    "Test Category",
			Description: "Test Description",
			Condition:   auction_entity.New,
			Status:      auction_entity.Active,
			Timestamp:   time.Now().Add(-30 * time.Second).Unix(),
		}

		_, err := collection.InsertOne(ctx, expiredAuction)
		if err != nil {
			t.Fatalf("Failed to insert expired auction: %v", err)
		}

		err = ar.CheckAuctionIsClosed(ctx, filter, duration)

		expectedError := &internal_error.InternalError{Message: "Auction has expired", Err: "internal_server_err"}
		assert.Equal(t, expectedError, err)
	})

	t.Run("Non-Expired Auction", func(t *testing.T) {
		notExpiredAuction := &auction_entity.Auction{
			Id:          "124",
			ProductName: "Another Product",
			Category:    "Another Category",
			Description: "Another Description",
			Condition:   auction_entity.New,
			Status:      auction_entity.Active,
			Timestamp:   time.Now().Add(30 * time.Second),
		}

		err := ar.CreateAuction(ctx, notExpiredAuction)
		if err != nil {
			t.Fatalf("Failed to create non-expired auction: %v", err)
		}

		err = ar.CheckAuctionIsClosed(ctx, bson.M{"_id": "124"}, duration)

		assert.Nil(t, err)
	})
}
