package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/muhmouddd21/rssAggregator/internal/db"
)

type apiConfig struct {
	DB *db.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("portString is not found in env")
	}
	dburl := os.Getenv("DB_URL")

	if dburl == "" {
		log.Fatal("dburl is not found in env")
	}

	conn, err := sql.Open("postgres", dburl)
	if err != nil {
		log.Fatal("can not connect to DB", err)
	}

	apicfg := apiConfig{
		DB: db.New(conn),
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handle_health)
	v1Router.Get("/err", handle_err)
	v1Router.Post("/users", apicfg.handlerCreateUser)
	v1Router.Get("/users", withLogging(apicfg.authedUser(apicfg.handleGetUser)))
	v1Router.Post("/feeds", apicfg.authedUser(apicfg.handlerCreateFeed))
	v1Router.Post("/feeds_follows", apicfg.authedUser(apicfg.handlerFollowFeed))
	v1Router.Get("/feeds_follows", withLogging(apicfg.authedUser(apicfg.handlerGetFollowsfeed)))
	v1Router.Delete("/feeds_follows/{feedfollowid}", withLogging(apicfg.authedUser(apicfg.handlerDeleteFollowsfeed)))
	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Addr:    ":" + portString,
		Handler: router,
	}
	fmt.Printf("server is running on %v", portString)

	err2 := srv.ListenAndServe()

	if err2 != nil {
		log.Fatal(err2)
	}
}
