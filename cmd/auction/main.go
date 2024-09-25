package main

import (
	"context"
	"github.com/joho/godotenv"
	"leilao/configuration/database/mongodb"
	"log"
)

func main() {
	ctx := context.Background()

	// godotenv.Load() to load the env, maybe you need to specify the path of .env
	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	_, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
