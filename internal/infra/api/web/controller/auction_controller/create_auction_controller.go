package auction_controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"leilao/configuration/rest_err"
	"leilao/internal/infra/api/web/validation"
	"leilao/internal/usecase/auction_usecase"
	"net/http"
)

type auctionController struct {
	auctionUseCase auction_usecase.AuctionUseCase
}

func NewAuctionController(auctionUseCase auction_usecase.AuctionUseCase) *auctionController {
	return &auctionController{
		auctionUseCase: auctionUseCase,
	}
}

func (ac *auctionController) CreateAuction(c *gin.Context) {
	var auctionInputDTO auction_usecase.AuctionInputDTO

	if err := c.ShouldBindJSON(&auctionInputDTO); err != nil {
		restErr := validation.ValidateErr(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	err := ac.auctionUseCase.CreateAuction(context.Background(), auctionInputDTO)
	if err != nil {
		restErr := rest_err.ConvertError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	c.Status(http.StatusCreated)
}
