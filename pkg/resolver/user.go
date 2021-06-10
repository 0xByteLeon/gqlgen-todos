package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/leon/gqlgen-todos/graph/model"
	pkgModel "github.com/leon/gqlgen-todos/pkg/model"
)

func (r *mutationResolver) Register(ctx context.Context, userName string) (*model.User, error) {
	return r.UserRepo.Register(ctx, userName)
}

func (r *queryResolver) QueryTodosByUserID(ctx context.Context, userID string) ([]*pkgModel.Todo, error) {
	user, err := r.UserRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	todos, err := r.TodoRepo.QueryTodos(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return todos, nil
}
