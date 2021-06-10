package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/leon/gqlgen-todos/graph/model"
)

func (r *mutationResolver) Register(ctx context.Context, userName string) (*model.User, error) {
	return r.UserRepo.Register(ctx, userName)
}

func (r *queryResolver) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	return r.UserRepo.GetUserByID(ctx, userID)
}
