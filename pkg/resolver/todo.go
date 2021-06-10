package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	model1 "github.com/leon/gqlgen-todos/graph/model"
	"github.com/leon/gqlgen-todos/pkg/generated"
	"github.com/leon/gqlgen-todos/pkg/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model1.NewTodo) (*model1.NewTodoResponse, error) {
	_, err := r.UserRepo.GetUserByID(ctx, input.UserID)
	if err != nil {
		return &model1.NewTodoResponse{
			Error: marshalError(err),
		}, nil
	}
	todo, err := r.TodoRepo.CreateTodo(ctx, input)
	return &model1.NewTodoResponse{
		Error: nil,
		Todo:  todo,
	}, nil
}

func (r *queryResolver) QueryTodosByUserID(ctx context.Context, userID string) ([]*model.Todo, error) {
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

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model1.User, error) {
	return r.UserRepo.GetUserByID(ctx, obj.UserID)
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
