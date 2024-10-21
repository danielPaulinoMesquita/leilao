package bid_usecase

import (
	"context"
	"leilao/internal/internal_error"
)

func (bu *BidUseCase) FindBidByAuctionId(
	ctx context.Context, auctionId string) ([]BidOutputDTO, *internal_error.InternalError) {
	bidList, err := bu.BidRepository.FindBidByAuctionId(ctx, auctionId)
	if err != nil {
		return nil, err
	}

	var bidListDTO []BidOutputDTO
	for _, bid := range bidList {
		bidListDTO = append(bidListDTO, BidOutputDTO{
			Id:        bid.Id,
			UserID:    bid.UserId,
			AuctionID: bid.AuctionId,
			Amount:    bid.Amount,
			Timestamp: bid.Timestamp,
		})
	}

	return bidListDTO, nil
}
func (bu *BidUseCase) FindWinningBidByAuctionId(
	ctx context.Context,
	auctionId string) (*BidOutputDTO, *internal_error.InternalError) {
	bid, err := bu.BidRepository.FindWinningBidByAuctionId(ctx, auctionId)
	if err != nil {
		return nil, err
	}

	bidOutPut := &BidOutputDTO{
		Id:        bid.Id,
		UserID:    bid.UserId,
		AuctionID: bid.AuctionId,
		Amount:    bid.Amount,
		Timestamp: bid.Timestamp,
	}

	return bidOutPut, nil
}
