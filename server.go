package main

import (
	"graphqlProject/database"
	"graphqlProject/graph"
	"graphqlProject/repo"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := database.Connection()
	if err != nil {
		log.Println("db connection %w", err)
		return
	}

	// dsn := "host=localhost user=postgres password=admin dbname=graphqldb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	panic(err)
	// }

	database, err := repo.NewConn(db)

	if err != nil {
		log.Fatalln(err)
	}

	dbb := repo.NewStore(database)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{S: dbb}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
