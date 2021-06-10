package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/google/uuid"
	graph "github.com/leon/gqlgen-todos/graph/model"
	"github.com/leon/gqlgen-todos/pkg/generated"
	pkg "github.com/leon/gqlgen-todos/pkg/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input graph.NewTodo) (*pkg.Todo, error) {
	user, err := r.UserRepo.GetUserByID(ctx, input.UserID)
	if err != nil {
		return nil, err
	}
	todo := pkg.Todo{
		ID:     uuid.New().String(),
		Text:   input.Text,
		Done:   false,
		UserID: user.ID,
	}
	return &todo, nil
}

func (r *queryResolver) Empty(ctx context.Context) (string, error) {
	return "", nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
