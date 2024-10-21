package bid

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"leilao/configuration/logger"
	"leilao/internal/entity/auction_entity"
	"leilao/internal/entity/bid_entity"
	"leilao/internal/infra/database/auction"
	"leilao/internal/internal_error"
	"sync"
)

type BidEntityMongo struct {
	Id        string  `bson:"id"`
	UserId    string  `bson:"user_id"`
	AuctionId string  `bson:"auction_id"`
	Amount    float64 `bson:"amount"`
	Timestamp int64   `bson:"timestamp"`
}

type BidRepository struct {
	Collection        *mongo.Collection
	AuctionRepository *auction.AuctionRepository
}

func NewBidRepository(database *mongo.Database, auctionRepository *auction.AuctionRepository) *BidRepository {
	return &BidRepository{
		Collection:        database.Collection("bid"),
		AuctionRepository: auctionRepository,
	}
}

func (bd *BidRepository) CreateBid(ctx context.Context, bidEntities []bid_entity.Bid) *internal_error.InternalError {
	var wq sync.WaitGroup

	for _, bid := range bidEntities {

		wq.Add(1)

		go func(bidValue bid_entity.Bid) {
			defer wq.Done()

			auctionEnity, err := bd.AuctionRepository.FindAuctionByID(ctx, bidValue.AuctionId)
			if err != nil {
				logger.Error("Error trying to find auction by id", err)
				return
			}

			if auctionEnity.Status != auction_entity.Active {
				return
			}

			bidEntityMongo := &BidEntityMongo{
				Id:        bidValue.Id,
				UserId:    bidValue.UserId,
				AuctionId: bidValue.AuctionId,
				Amount:    bidValue.Amount,
				Timestamp: bidValue.Timestamp.Unix(),
			}

			if _, err := bd.Collection.InsertOne(ctx, bidEntityMongo); err != nil {
				logger.Error("Error trying to insert bid", err)
				return
			}

		}(bid)

	}

	wq.Wait()
	return nil

}
