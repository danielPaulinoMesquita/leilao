package user

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"leilao/configuration/logger"
	"leilao/internal/entity/user_entity"
	"leilao/internal/internal_error"
)

type UserEntityMongo struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
}

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database) *UserRepository {
	return &UserRepository{
		Collection: database.Collection("users"),
	}
}

func (repo *UserRepository) FindById(ctx context.Context,
	userId string) (*user_entity.User, *internal_error.InternalError) {
	filter := bson.M{"_id": userId}

	var user UserEntityMongo
	err := repo.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(fmt.Sprintf("User not found with this id = %d", userId), err)
			return nil, internal_error.NewNotFoundError(
				fmt.Sprintf("User not found with this id = %d", userId))
		}

		logger.Error("Error trying to find user by userId", err)
		return nil, internal_error.NewNotFoundError("Error trying to find user by userId")
	}

	userEntity := &user_entity.User{
		Id:   user.Id,
		Name: user.Name,
	}

	return userEntity, nil
}
