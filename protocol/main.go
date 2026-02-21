//go:generate ./generate_bindings.sh

package main

import (
	"context"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/optimus-boc-protocol/services"
	"github.com/optimus-boc-protocol/services/did"
	ctl "github.com/optimus-boc-protocol/controllers/did"
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

	// initialize blockchain services
	rpc := os.Getenv("CHAIN_RPC_URL")
	if rpc == "" {
		log.Fatal("CHAIN_RPC_URL is required")
	}

	bnplAddr := os.Getenv("BNPL_MANAGER_ADDRESS")
	if bnplAddr == "" {
		log.Fatal("BNPL_MANAGER_ADDRESS is required")
	}

	bnplSvc, err := services.NewBNPLService(rpc, bnplAddr)
	if err != nil {
		log.Fatalf("failed to create bnpl service: %v", err)
	}

	// DID service/controller
	didAddr := os.Getenv("DID_REGISTRY_ADDRESS")
	if didAddr == "" {
		log.Fatal("DID_REGISTRY_ADDRESS is required")
	}
	didSvc, err := did.NewDid(rpc, didAddr)
	if err != nil {
		log.Fatalf("failed to create did service: %v", err)
	}
	didCtrl := ctl.NewController(didSvc)
	r.Mount("/", didCtrl.Routes())

	// basic route demonstrating bnpl service usage
	r.Get("/arrangement/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id := new(big.Int)
		id.SetString(idStr, 10)
		arr, err := bnplSvc.GetArrangement(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(arr)
	})

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