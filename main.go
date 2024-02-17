package main

import (
	"blog/internal/database"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	db *database.Queries
}

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("can't connect to database...")
	}

	connection, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("can't connect to database...", err)
	}

	queries := database.New(connection)
	apiCfg := apiConfig{
		db: queries,
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	v1Router := chi.NewRouter()
	v1Router.Get("/ready", handlerReadiness)
	v1Router.Get("/notready", handlerError)
	v1Router.Post("/users", apiCfg.handlerUserCreate)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlergetUser))
	r.Mount("/v1", v1Router)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// start the server
	fmt.Printf("Server starting on port %s .... \n", port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
