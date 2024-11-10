package auction_entity

import (
	"context"
	"github.com/google/uuid"
	"leilao/internal/internal_error"
	"time"
)

func CreateAuction(
	productName string, category, description string,
	condition ProductionCondition) (*Auction, *internal_error.InternalError) {

	auction := &Auction{
		Id:          uuid.New().String(),
		ProductName: productName,
		Category:    category,
		Description: description,
		Condition:   condition,
		Status:      Active,
		Timestamp:   time.Now(),
	}

	if err := auction.Validate(); err != nil {
		return nil, err
	}

	return auction, nil
}

func (au *Auction) Validate() *internal_error.InternalError {
	if len(au.ProductName) <= 1 {
		return internal_error.NewBadRequestError("invalid product name")
	}
	if len(au.Category) <= 2 {
		return internal_error.NewBadRequestError("invalid category")
	}
	if len(au.Description) <= 10 {
		return internal_error.NewBadRequestError("invalid description")
	}
	return nil
}

type Auction struct {
	Id          string
	ProductName string
	Category    string
	Description string
	Condition   ProductionCondition
	Status      AuctionStatus
	Timestamp   time.Time
}

type ProductionCondition int
type AuctionStatus int

const (
	Active AuctionStatus = iota
	Completed
	Closed
)

const (
	New ProductionCondition = iota
	Us
	Refurbished
)

type AuctionRepositoryInterface interface {
	CreateAuction(ctx context.Context,
		auctionEntity *Auction) *internal_error.InternalError

	FindAuctions(
		ctx context.Context,
		status AuctionStatus,
		category, productName string,
	) ([]Auction, *internal_error.InternalError)

	FindAuctionByID(
		ctx context.Context,
		id string) (*Auction, *internal_error.InternalError)
}
