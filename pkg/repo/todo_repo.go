package repo

import (
	"context"

	"github.com/leon/gqlgen-todos/graph/model"
	pkgModel "github.com/leon/gqlgen-todos/pkg/model"
)

type TodoRepo interface {
	CreateTodo(ctx context.Context, input model.NewTodo) (*pkgModel.Todo, error)
	QueryTodos(ctx context.Context, userID string) ([]*pkgModel.Todo, error)
}

func NewTodoRepo() TodoRepo {
	return newTodoDB()
}
