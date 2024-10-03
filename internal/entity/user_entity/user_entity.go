package user_entity

import (
	"context"
	"leilao/internal/internal_error"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserRepositoryInterface interface {
	FindById(ctx context.Context,
		userId string) (*User, *internal_error.InternalError)
}
