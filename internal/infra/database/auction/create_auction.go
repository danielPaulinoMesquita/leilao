package auction

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"leilao/configuration/logger"
	"leilao/internal/entity/auction_entity"
	"leilao/internal/internal_error"
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

func (ar *AuctionRepository) CreateAuction(ctx context.Context, auction *auction_entity.Auction) *internal_error.InternalError {
	auctionEntity := &AuctionEntityMongo{
		Id:          auction.Id,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   auction.Condition,
		Status:      auction.Status,
		Timestamp:   auction.Timestamp.Unix(),
	}

	_, err := ar.Collection.InsertOne(ctx, auctionEntity)
	if err != nil {
		logger.Error("Error trying to insert", err)
		return internal_error.NewInternalServerError("Error trying to insert auction in database")
	}

	return nil
}
