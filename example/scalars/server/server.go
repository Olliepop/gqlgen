package main

import (
	"log"
	"net/http"

	"github.com/vektah/gqlgen/example/scalars"
	"github.com/vektah/gqlgen/handler"
)

func main() {
	http.Handle("/", handler.GraphiQL("Starwars", "/query"))
	http.Handle("/query", handler.GraphQL(scalars.NewExecutor(&scalars.Resolver{})))

	log.Fatal(http.ListenAndServe(":8084", nil))
}
