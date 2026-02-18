package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	r := chi.NewRouter()
	r.Get("/health", func(w http.ResponseWriter, r * http.Request){
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	r.Get("/", func(w http.ResponseWriter, r * http.Request){
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, Optimus Protocol"))
	})

	port := os.Getenv("PORT")
	if port == "" { port = "8000"}

	dbUrl := os.Getenv("DATABASE_URL")
	pool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil { log.Fatalf("db connect: %v", err) }
	defer pool.Close()

	srv := &http.Server{
		Addr:	":" + port,
		Handler: r,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("server listening on", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error %v", err)
	}
}