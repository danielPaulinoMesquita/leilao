package bid_usecase

import (
	"leilao/internal/entity/bid_entity"
	"time"
)

type BidOutputDTO struct {
	Id        string    `json:"id"`
	UserID    string    `json:"user_id"`
	AuctionID string    `json:"auction_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

type BidUseCase struct {
	BidRepository bid_entity.BidEntityRepositoryInterface
}
