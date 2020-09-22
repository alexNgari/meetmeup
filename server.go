package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-pg/pg/v10"
	"github.com/alexNgari/meetmeup/graph"
	"github.com/alexNgari/meetmeup/graph/generated"
	"github.com/alexNgari/meetmeup/postgres"
)

const defaultPort = "8080"

func main() {

	DB := postgres.New(&pg.Options{
		User:"meetmeup_user",
		Password:"meetmeup",
		Database:"meetmeup_dev",
	})

	defer DB.Close()

	DB.AddQueryHook(postgres.DBLogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		MeetupsRepo: postgres.MeetupsRepo{DB: DB},
		UsersRepo: postgres.UsersRepo{DB: DB},
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
