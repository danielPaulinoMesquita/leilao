package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"leilao/configuration/database/mongodb"
	"leilao/internal/infra/api/web/controller/auction_controller"
	"leilao/internal/infra/api/web/controller/bid_controller"
	"leilao/internal/infra/api/web/controller/user_controller"
	"leilao/internal/infra/database/auction"
	"leilao/internal/infra/database/bid"
	"leilao/internal/infra/database/user"
	"leilao/internal/usecase/auction_usecase"
	"leilao/internal/usecase/bid_usecase"
	"leilao/internal/usecase/user_usecase"
	"log"
)

func main() {
	ctx := context.Background()

	// godotenv.Load() to load the env, maybe you need to specify the path of .env
	//if err := godotenv.Load(".env"); err != nil { to run local
	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return
	}

	databaseConnection, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	router := gin.Default()

	userController, bidController, auctionsController := initDependencies(databaseConnection)

	router.GET("/auctions", auctionsController.FindAuctions)
	router.GET("/auctions/:auctionId", auctionsController.FindAuctionById)
	router.POST("/auctions", auctionsController.CreateAuction)
	router.GET("/auctions/winner/:auctionId", auctionsController.FindWinningBidByAuctionId)
	router.POST("/bid", bidController.CreateBid)
	router.GET("/bid/:auctionId", bidController.FindBidByAuctionId)
	router.GET("/user/:userId", userController.FindUserById)

	router.Run(":8080")

}

func initDependencies(database *mongo.Database) (
	userController *user_controller.UserController,
	bidController *bid_controller.BidController,
	auctionController *auction_controller.AuctionController) {

	auctionRespository := auction.NewAuctionRepository(database)
	bidRespository := bid.NewBidRepository(database, auctionRespository)
	userRespository := user.NewUserRepository(database)

	userController = user_controller.NewUserController(user_usecase.NewUserUseCase(userRespository))
	auctionController = auction_controller.NewAuctionController(auction_usecase.NewAuctionUseCase(auctionRespository, bidRespository))
	bidController = bid_controller.NewBidController(bid_usecase.NewBidUseCase(bidRespository))

	return
}
