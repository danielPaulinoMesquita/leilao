package auction_usecase

import (
	"context"
	"leilao/internal/entity/auction_entity"
	"leilao/internal/entity/bid_entity"
	"leilao/internal/infra/database/bid"
	"leilao/internal/internal_error"
	"leilao/internal/usecase/bid_usecase"
	"time"
)

type AuctionInputDTO struct {
	ProductName string              `json:"product_name" binding:"required,min=1"`
	Category    string              `json:"category" binding:"required,min=2"`
	Description string              `json:"description" binding:"required,min=10, max=200"`
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

type WinningInfoOutputDTO struct {
	Auction AuctionOutputDTO          `json:"auction"`
	Bid     *bid_usecase.BidOutputDTO `json:"bid,omitempty"`
}
type ProductionCondition int64
type AuctionStatus int64

func NewAuctionUseCase(auctionRepositoryInterface auction_entity.AuctionRepositoryInterface,
	bidRepositoryInterface *bid.BidRepository) AuctionUseCaseInterface {
	return &AuctionUseCase{
		auctionRepository: auctionRepositoryInterface,
		bidRepository:     bidRepositoryInterface}
}

type AuctionUseCaseInterface interface {
	CreateAuction(
		ctx context.Context,
		auctionInput AuctionInputDTO) *internal_error.InternalError

	FindAuctions(
		ctx context.Context,
		status auction_entity.AuctionStatus,
		category, productName string) ([]AuctionOutputDTO, *internal_error.InternalError)

	FindAuctionById(
		ctx context.Context,
		id string) (*AuctionOutputDTO, *internal_error.InternalError)

	FindWinningBidByAuctionId(ctx context.Context,
		auctionId string) (*WinningInfoOutputDTO, *internal_error.InternalError)
}

type AuctionUseCase struct {
	auctionRepository auction_entity.AuctionRepositoryInterface
	bidRepository     bid_entity.BidEntityRepositoryInterface
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
