//go:generate go run ../../scripts/gqlgen.go

package resolver

import (
	"strconv"
	"strings"

	"github.com/leon/gqlgen-todos/graph/model"
	"github.com/leon/gqlgen-todos/pkg/repo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserRepo repo.UserRepo
	TodoRepo repo.TodoRepo
}

func marshalError(e error) *model.Error {
	if e == nil {
		return nil
	}
	errstr := e.Error()
	// add more log
	l := strings.Split(errstr, ":")
	if len(l) < 2 {
		return &model.Error{
			Code:    0,
			Message: errstr,
		}
	}
	code, err := strconv.ParseInt(l[0], 10, 64)
	if err != nil {
		return &model.Error{
			Code:    0,
			Message: errstr,
		}
	}
	return &model.Error{
		Code:    int(code),
		Message: strings.Join(l[1:], ""),
	}
}
