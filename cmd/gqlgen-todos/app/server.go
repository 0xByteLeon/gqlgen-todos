package app

import (
	"log"
	"net/http"
	"os"

	"github.com/leon/gqlgen-todos/pkg/repo"
	"github.com/leon/gqlgen-todos/pkg/resolver"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/leon/gqlgen-todos/pkg/generated"
)

const defaultPort = "8080"

type App struct {
}

func New() *App {
	return &App{}
}

func (a *App) Start() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{
		UserRepo: repo.NewUserRepo(),
		TodoRepo: repo.NewTodoRepo(),
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
