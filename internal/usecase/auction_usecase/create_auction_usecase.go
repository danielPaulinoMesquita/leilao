package auction_usecase

import (
	"context"
	"leilao/internal/entity/auction_entity"
	"leilao/internal/internal_error"
	"time"
)

type AuctionInputDTO struct {
	ProductName string              `json:"product_name"`
	Category    string              `json:"category"`
	Description string              `json:"description"`
	Condition   ProductionCondition `json:"condition"`
}

type AuctionOutputDTO struct {
	Id          string              `json:"id"`
	ProductName string              `json:"product_name"`
	Category    string              `json:"category"`
	Description string              `json:"description"`
	Condition   ProductionCondition `json:"condition"`
	Status      AuctionStatus       `json:"status"`
	Timestamp   time.Time           `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

type ProductionCondition int64
type AuctionStatus int64

type AuctionUseCase struct {
	auctionRepository auction_entity.AuctionRepositoryInterface
}

func (au *AuctionUseCase) CreateAuction(ctx context.Context, auctionInput AuctionInputDTO) *internal_error.InternalError {
	auction, err := auction_entity.CreateAuction(
		auctionInput.ProductName,
		auctionInput.Category,
		auctionInput.Description,
		auction_entity.ProductionCondition(auctionInput.Condition))

	if err != nil {
		return err
	}

	if err := au.auctionRepository.CreateAuction(ctx, auction); err != nil {
		return err
	}

	return nil
}
