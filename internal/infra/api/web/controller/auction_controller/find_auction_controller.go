package auction_controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"leilao/configuration/rest_err"
	"leilao/internal/entity/auction_entity"
	"net/http"
	"strconv"
)

func (ac *AuctionController) FindAuctionById(c *gin.Context) {
	auctionId := c.Param("auctionId")

	if err := uuid.Validate(auctionId); err != nil {
		errReset := rest_err.NewBadRequestError("Invalid fields", rest_err.Causes{
			Field:   "auctionId",
			Message: "Invalid UUID fields",
		})

		c.JSON(errReset.Code, errReset)
		return
	}

	auctionData, err := ac.auctionUseCase.FindAuctionById(context.Background(), auctionId)
	if err != nil {
		restErr := rest_err.ConvertError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, auctionData)
}

func (ac *AuctionController) FindAuctions(c *gin.Context) {
	status := c.Query("status")
	category := c.Query("category")
	productName := c.Query("productName")

	statusNumber, errConv := strconv.Atoi(status)
	if errConv != nil {
		errRest := rest_err.NewBadRequestError("Error trying to validate auction status param")
		c.JSON(errRest.Code, errRest)
		return
	}

	auctions, err := ac.auctionUseCase.FindAuctions(context.Background(),
		auction_entity.AuctionStatus(statusNumber), category, productName)

	if err != nil {
		restErr := rest_err.ConvertError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, auctions)
}

func (ac *AuctionController) FindWinningBidByAuctionId(c *gin.Context) {
	auctionId := c.Query("auctionId")

	if err := uuid.Validate(auctionId); err != nil {
		errReset := rest_err.NewBadRequestError("Invalid fields", rest_err.Causes{
			Field:   "auctionId",
			Message: "Invalid UUID fields",
		})

		c.JSON(errReset.Code, errReset)
		return
	}

	auctionData, err := ac.auctionUseCase.FindWinningBidByAuctionId(context.Background(), auctionId)
	if err != nil {
		restErr := rest_err.ConvertError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, auctionData)
}
