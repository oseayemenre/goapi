package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/oseayemenre/goapi/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}

	db_url := os.Getenv("DB_URL")
	if db_url == "" {
		log.Fatal("SQL is not found in the environment")
	}

	conn, err := sql.Open("postgres", db_url)

	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}

	apiCfg := &apiConfig{
		DB: database.New(conn),
	}

	r := chi.NewRouter()
	
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*", "https://*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	tr := chi.NewRouter()
	tr.Get("/", handlerTest)
	tr.Get("/err", handlerTestError)
	r.Post("/user", apiCfg.handlerUser)

	r.Mount("/health", tr)

	s := http.Server{
		Addr: ":" + port,
		Handler: r,
	}

	log.Printf("Server starting on port %v", port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}