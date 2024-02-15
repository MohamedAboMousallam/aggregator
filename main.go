package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	v1Router := chi.NewRouter()
	v1Router.Get("/ready", handlerReadiness)
	v1Router.Get("/notready", handlerError)
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
