package auction

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"leilao/configuration/logger"
	"leilao/internal/entity/auction_entity"
	"leilao/internal/internal_error"
	"os"
	"time"
)

type AuctionEntityMongo struct {
	Id          string                             `bson:"_id"`
	ProductName string                             `bson:"product_name"`
	Category    string                             `bson:"category"`
	Description string                             `bson:"description"`
	Condition   auction_entity.ProductionCondition `bson:"condition"`
	Status      auction_entity.AuctionStatus       `bson:"status"`
	Timestamp   int64                              `bson:"timestamp"`
}

type AuctionRepository struct {
	Collection *mongo.Collection
}

func NewAuctionRepository(database *mongo.Database) *AuctionRepository {
	return &AuctionRepository{
		Collection: database.Collection("auctions"),
	}
}

func GetAuctionDuration() (time.Duration, error) {
	auctionDurationStr := os.Getenv("AUCTION_DURATION")
	if auctionDurationStr == "" {
		return 0, fmt.Errorf("AUCTION_DURATION is not set")
	}

	duration, err := time.ParseDuration(auctionDurationStr)
	if err != nil {
		return 0, fmt.Errorf("invalid AUCTION_DURATION: %v", err)
	}

	return duration, nil
}

func (ar *AuctionRepository) CheckAuctionIsClosed(
	ctx context.Context,
	filter bson.M,
	duration time.Duration) *internal_error.InternalError {
	var auctionEntityMongo AuctionEntityMongo

	now := time.Now()
	expirationTime := time.Unix(auctionEntityMongo.Timestamp, 0).Add(duration)
	if now.After(expirationTime) {
		update := bson.M{"$set": bson.M{"status": auction_entity.Closed}}
		if _, err := ar.Collection.UpdateOne(ctx, filter, update); err != nil {
			logger.Error("Error closing expired auction", err)
		}
		return internal_error.NewInternalServerError("Auction has expired")
	}

	return nil
}

func (ar *AuctionRepository) CreateAuction(
	ctx context.Context,
	auction *auction_entity.Auction) *internal_error.InternalError {

	auctionEntity := &AuctionEntityMongo{
		Id:          auction.Id,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   auction.Condition,
		Status:      auction.Status,
		Timestamp:   auction.Timestamp.Unix(),
	}

	println("Inserting Data")
	fmt.Printf("Auction entity to insert: %+v\n", auctionEntity)

	_, err := ar.Collection.InsertOne(ctx, auctionEntity)
	if err != nil {
		logger.Error("Error trying to insert", err)
		fmt.Print("Error: ", err.Error())
		return internal_error.NewInternalServerError("Error trying to insert auction in database")
	}

	return nil
}
