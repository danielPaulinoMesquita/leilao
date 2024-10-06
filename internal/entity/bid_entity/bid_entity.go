package bid_entity

import (
	"context"
	"leilao/internal/internal_error"
	"time"
)

type Bid struct {
	Id        string
	UserId    string
	AuctionId string
	Amount    float64
	Timestamp time.Time
}

type BidEntityRepositoryInterface interface {
	CreateBid(ctx context.Context, bidEntities []Bid) *internal_error.InternalError
	FindBidByAuctionId(
		auctionId string, ctx context.Context) ([]Bid, *internal_error.InternalError)
	FindWinningBidByAuctionId(
		ctx context.Context,
		auctionId string) (*Bid, *internal_error.InternalError)
}
