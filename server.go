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
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
	customMiddleware "github.com/alexNgari/meetmeup/middleware"
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

	userRepo := postgres.UsersRepo{DB: DB},

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		Debug: true,
		AllowCredentials: true,
		AllowedOrigins: []string{"http://localhost:8080"}
	}).Handler)

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(customMiddleware.AuthMiddleware(userRepo))

	c := generated.Config{Resolvers: &graph.Resolver{
				MeetupsRepo: postgres.MeetupsRepo{DB: DB},
				UsersRepo: userRepo
			}}

	queryHandler := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", graph.DataloaderMiddleware(DB, queryHandler))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
