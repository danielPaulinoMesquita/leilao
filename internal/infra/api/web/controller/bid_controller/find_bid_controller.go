package bid_controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"leilao/configuration/rest_err"
	"net/http"
)

func (bd *BidController) FindBidByAuctionId(c *gin.Context) {
	auctionId := c.Query("auctionId")

	if err := uuid.Validate(auctionId); err != nil {
		errReset := rest_err.NewBadRequestError("Invalid fields", rest_err.Causes{
			Field:   "auctionId",
			Message: "Invalid UUID fields",
		})

		c.JSON(errReset.Code, errReset)
		return
	}

	bidOuputList, err := bd.bidUseCase.FindBidByAuctionId(context.Background(), auctionId)
	if err != nil {
		restErr := rest_err.ConvertError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, bidOuputList)
}
