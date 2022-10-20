package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	auth "github.com/jbutcher93/quotes-starter/Auth"
	"github.com/jbutcher93/quotes-starter/gqlgen/graph"
	"github.com/jbutcher93/quotes-starter/gqlgen/graph/generated"
)

const defaultPort = "8081"

func main() {
	router := chi.NewRouter()
	router.Use(auth.Middleware())
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
