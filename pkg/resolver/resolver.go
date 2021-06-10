//go:generate go run ../../scripts/gqlgen.go

package resolver

import (
	"github.com/leon/gqlgen-todos/pkg/repo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserRepo repo.UserRepo
	TodoRepo repo.TodoRepo
}
