package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	graphql_todo_practice "github.com/hotdog132/graphql-todo-practice"
	"github.com/hotdog132/graphql-todo-practice/config"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	configs, err := config.NewConfigurations("development")
	if err != nil {
		log.Println("Get configs error:", err)
	}

	graphql_todo_practice.Configs = configs

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(graphql_todo_practice.NewExecutableSchema(graphql_todo_practice.Config{Resolvers: &graphql_todo_practice.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
