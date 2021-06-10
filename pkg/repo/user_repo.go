package repo

import (
	"context"

	"github.com/leon/gqlgen-todos/graph/model"
)

type UserRepo interface {
	GetUserByID(ctx context.Context, userID string) (*model.User, error)
	Register(ctx context.Context, userName string) (*model.User, error)
}

func NewUserRepo() UserRepo {
	return newUserDB()
}
