package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	model1 "github.com/leon/gqlgen-todos/graph/model"
	"github.com/leon/gqlgen-todos/pkg/generated"
	"github.com/leon/gqlgen-todos/pkg/model"
)

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model1.User, error) {
	return r.UserRepo.GetUserByID(ctx, obj.UserID)
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
