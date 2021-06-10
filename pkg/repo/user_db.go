package repo

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/leon/gqlgen-todos/graph/model"
)

type UserDB struct {
	users []*model.User
}

func (db *UserDB) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	for _, user := range db.users {
		if user.ID == userID {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (db *UserDB) Register(ctx context.Context, userName string) (*model.User, error) {
	for _, user := range db.users {
		if user.Name == userName {
			return nil, errors.New("")
		}
	}
	newUser := &model.User{
		ID:   uuid.New().String(),
		Name: userName,
	}
	db.users = append(db.users, newUser)
	return newUser, nil
}

func newUserDB() UserRepo {
	return &UserDB{}
}
