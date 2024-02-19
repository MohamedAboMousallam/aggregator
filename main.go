package main

import (
	"blog/internal/database"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	db *database.Queries
}

func main() {

	feed, errror := urlToFeed("https://blog.boot.dev/index.xml")
	if errror != nil {
		log.Fatal(errror)
	}
	fmt.Println(feed)

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

	const collectionConcurrency = 10
	const collectionInterval = time.Minute

	go startScraping(queries, collectionConcurrency, collectionInterval)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	v1Router := chi.NewRouter()

	v1Router.Get("/ready", handlerReadiness)
	v1Router.Get("/notready", handlerError)

	v1Router.Post("/users", apiCfg.handlerUserCreate)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlergetUser))

	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)

	v1Router.Post("/feeds_follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollows))
	v1Router.Get("/feeds_follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))
	v1Router.Delete("/feeds_follows/{feedfollowsID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollows))

	v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handlerPostsGet))

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
