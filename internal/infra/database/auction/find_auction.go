package auction

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"leilao/configuration/logger"
	"leilao/internal/entity/auction_entity"
	"leilao/internal/internal_error"
	"time"
)

func (ar *AuctionRepository) FindAuctionByID(ctx context.Context, id string) (*auction_entity.Auction, *internal_error.InternalError) {
	filter := bson.M{"_id": id}

	var auctionEntityMongo AuctionEntityMongo
	if err := ar.Collection.FindOne(ctx, filter).Decode(&auctionEntityMongo); err != nil {
		logger.Error(fmt.Sprintf("Error trying to find auction by id = %s", id), err)
		return nil, internal_error.NewInternalServerError("Error trying to find auction")
	}

	duration, err := GetAuctionDuration()
	if err != nil {
		logger.Error("Error getting auction duration", err)
		return nil, internal_error.NewInternalServerError("Failed to get auction duration")
	}

	err = ar.CheckAuctionIsClosed(ctx, filter, duration)

	if err != nil {
		return nil, internal_error.NewInternalServerError(err.Error())
	}

	return &auction_entity.Auction{
		Id:          auctionEntityMongo.Id,
		ProductName: auctionEntityMongo.ProductName,
		Category:    auctionEntityMongo.Category,
		Description: auctionEntityMongo.Description,
		Condition:   auctionEntityMongo.Condition,
		Status:      auctionEntityMongo.Status,
		Timestamp:   time.Unix(auctionEntityMongo.Timestamp, 0),
	}, nil
}

func (ar *AuctionRepository) FindAuctions(
	ctx context.Context,
	status auction_entity.AuctionStatus,
	category string,
	productName string) ([]auction_entity.Auction, *internal_error.InternalError) {
	filter := bson.M{}

	if status != 0 {
		filter["status"] = status
	}

	if category != "" {
		filter["category"] = category
	}

	if productName != "" {
		filter["product_name"] = primitive.Regex{
			Pattern: productName,
			Options: "i",
		}
	}

	fmt.Println("Filter: ", filter)

	cursor, err := ar.Collection.Find(ctx, filter)

	if err != nil {
		logger.Error("Error while finding auctions", err)
		return nil, internal_error.NewInternalServerError("Error trying  to find an auction")
	}

	if cursor.RemainingBatchLength() == 0 {
		fmt.Println("No auctions found.")
	}

	defer cursor.Close(ctx)

	var auctionEntityMongo []AuctionEntityMongo
	if err := cursor.All(ctx, &auctionEntityMongo); err != nil {
		logger.Error("Error while finding auctions", err)
		return nil, internal_error.NewInternalServerError("Error trying to find an auction")
	}

	var auctionEntity []auction_entity.Auction
	for _, auctionMongo := range auctionEntityMongo {
		auctionEntity = append(auctionEntity, auction_entity.Auction{
			Id:          auctionMongo.Id,
			ProductName: auctionMongo.ProductName,
			Category:    auctionMongo.Category,
			Description: auctionMongo.Description,
			Condition:   auctionMongo.Condition,
			Status:      auctionMongo.Status,
			Timestamp:   time.Unix(auctionMongo.Timestamp, 0),
		})
	}

	return auctionEntity, nil
}
