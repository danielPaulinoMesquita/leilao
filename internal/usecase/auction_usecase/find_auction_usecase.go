package auction_usecase

import (
	"context"
	"leilao/internal/entity/auction_entity"
	"leilao/internal/internal_error"
)

func (au *AuctionUseCase) FindAuctions(ctx context.Context, status auction_entity.AuctionStatus, category, productName string) ([]AuctionOutputDTO, *internal_error.InternalError) {
	auctions, err := au.auctionRepository.FindAuctions(ctx, auction_entity.AuctionStatus(status), category, productName)

	if err != nil {
		return nil, err
	}

	var auctionsOutput []AuctionOutputDTO

	for _, auction := range auctions {
		auctionsOutput = append(auctionsOutput, AuctionOutputDTO{
			Id:          auction.Id,
			ProductName: auction.ProductName,
			Category:    auction.Category,
			Description: auction.Description,
			Condition:   ProductionCondition(auction.Condition),
			Status:      AuctionStatus(auction.Status),
			Timestamp:   auction.Timestamp,
		})
	}
	return auctionsOutput, nil

}

func (au *AuctionUseCase) FindAuctionById(ctx context.Context, id string) (*AuctionOutputDTO, *internal_error.InternalError) {
	auctionEntity, err := au.auctionRepository.FindAuctionByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &AuctionOutputDTO{
		Id:          auctionEntity.Id,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   ProductionCondition(auctionEntity.Condition),
		Status:      AuctionStatus(auctionEntity.Status),
		Timestamp:   auctionEntity.Timestamp,
	}, nil
}
