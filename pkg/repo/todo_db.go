package repo

import (
	"context"

	"github.com/google/uuid"

	"github.com/leon/gqlgen-todos/graph/model"
	pkgModel "github.com/leon/gqlgen-todos/pkg/model"
)

type TodoDB struct {
	todos []*pkgModel.Todo
}

func (db *TodoDB) CreateTodo(ctx context.Context, input model.NewTodo) (*pkgModel.Todo, error) {
	todo := pkgModel.Todo{
		ID:     uuid.New().String(),
		Text:   input.Text,
		Done:   false,
		UserID: input.UserID,
	}
	db.todos = append(db.todos, &todo)
	return &todo, nil
}

func (db *TodoDB) QueryTodos(ctx context.Context, userID string) ([]*pkgModel.Todo, error) {
	todoList := make([]*pkgModel.Todo, 0)
	for _, todo := range db.todos {
		if userID == todo.UserID {
			todoList = append(todoList, todo)
		}
	}
	return todoList, nil
}

func newTodoDB() TodoRepo {
	return &TodoDB{}
}
