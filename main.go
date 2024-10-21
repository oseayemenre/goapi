package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port:=os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment")
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