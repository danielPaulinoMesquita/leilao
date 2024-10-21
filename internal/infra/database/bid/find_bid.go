package bid

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"leilao/configuration/logger"
	"leilao/internal/entity/bid_entity"
	"leilao/internal/internal_error"
	"time"
)

func (bd *BidRepository) FindBidByAuctionId(
	ctx context.Context, auctionId string) ([]bid_entity.Bid, *internal_error.InternalError) {
	filter := bson.M{"auction_id": auctionId}

	cursor, err := bd.Collection.Find(ctx, filter)
	if err != nil {
		logger.Error(fmt.Sprintf("Error trying to find bids by auction %s", auctionId), err)
		return nil, internal_error.NewInternalServerError(
			fmt.Sprintf("Error trying to find bids by auction %s", auctionId))
	}

	var bidEntityMongo []BidEntityMongo
	if err = cursor.All(ctx, &bidEntityMongo); err != nil {
		logger.Error(fmt.Sprintf("Error trying to find bids by auction %s", auctionId), err)
		return nil, internal_error.NewInternalServerError(
			fmt.Sprintf("Error trying to find bids by auction %s", auctionId))

	}

	var bidEntites []bid_entity.Bid
	for _, b := range bidEntityMongo {
		bidEntites = append(bidEntites, bid_entity.Bid{
			Id:        b.Id,
			UserId:    b.UserId,
			AuctionId: b.AuctionId,
			Amount:    b.Amount,
			Timestamp: time.Unix(int64(b.Timestamp), 0),
		})
	}

	return bidEntites, nil
}

func (bd *BidRepository) FindWinningBidByAuctionId(
	ctx context.Context,
	auctionId string) (*bid_entity.Bid, *internal_error.InternalError) {
	filter := bson.M{"auction_id": auctionId}

	var bidEntityMongo BidEntityMongo
	opts := options.FindOne().SetSort(bson.D{{"amount", -1}})
	if err := bd.Collection.FindOne(ctx, filter, opts).Decode(&bidEntityMongo); err != nil {
		logger.Error("Error trying to find the auction winner", err)
		return nil, internal_error.NewInternalServerError("Error trying to find the auction winner")
	}

	return &bid_entity.Bid{
		Id:        bidEntityMongo.Id,
		UserId:    bidEntityMongo.UserId,
		AuctionId: bidEntityMongo.AuctionId,
		Amount:    bidEntityMongo.Amount,
		Timestamp: time.Unix(bidEntityMongo.Timestamp, 0),
	}, nil

}
