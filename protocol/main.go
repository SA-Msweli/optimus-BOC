//go:generate ./generate_bindings.sh

package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	// Service sub-packages (interface + implementation)
	bnplsvc "github.com/optimus-boc-protocol/services/bnpl"
	daosvc "github.com/optimus-boc-protocol/services/dao"
	"github.com/optimus-boc-protocol/services/did"
	loansvc "github.com/optimus-boc-protocol/services/loan"
	tvsvc "github.com/optimus-boc-protocol/services/tokenvault"

	// Controllers
	bnplctl "github.com/optimus-boc-protocol/controllers/bnpl"
	daoctl "github.com/optimus-boc-protocol/controllers/dao"
	didctl "github.com/optimus-boc-protocol/controllers/did"
	loanctl "github.com/optimus-boc-protocol/controllers/loan"
	tvctl "github.com/optimus-boc-protocol/controllers/tokenvault"

	"github.com/optimus-boc-protocol/eth"
	"github.com/optimus-boc-protocol/store"
)

func main() {
	cfg := LoadConfig()

	// ---------- database ----------
	pool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}
	defer pool.Close()
	dbStore := store.New(pool)

	// ---------- ethereum transactor ----------
	auth, _, err := eth.NewTransactor(context.Background(), cfg.ChainRPCURL)
	if err != nil {
		log.Fatalf("transactor: %v", err)
	}

	// ---------- services ----------
	didSvc, err := did.NewDid(cfg.ChainRPCURL, cfg.DIDRegistryAddr)
	if err != nil {
		log.Fatalf("did service: %v", err)
	}

	bnplSvc, err := bnplsvc.NewService(cfg.ChainRPCURL, cfg.BNPLManagerAddr)
	if err != nil {
		log.Fatalf("bnpl service: %v", err)
	}

	daoSvc, err := daosvc.NewService(cfg.ChainRPCURL, cfg.DAOManagerAddr)
	if err != nil {
		log.Fatalf("dao service: %v", err)
	}

	loanSvc, err := loansvc.NewService(cfg.ChainRPCURL, cfg.LoanManagerAddr)
	if err != nil {
		log.Fatalf("loan service: %v", err)
	}

	tvSvc, err := tvsvc.NewService(cfg.ChainRPCURL, cfg.TokenVaultAddr)
	if err != nil {
		log.Fatalf("tokenvault service: %v", err)
	}

	// ---------- controllers ----------
	didCtrl := didctl.NewController(didSvc, dbStore, auth)
	bnplCtrl := bnplctl.NewController(bnplSvc, dbStore, auth)
	daoCtrl := daoctl.NewController(daoSvc, dbStore, auth)
	loanCtrl := loanctl.NewController(loanSvc, dbStore, auth)
	tvCtrl := tvctl.NewController(tvSvc, auth)

	// ---------- router ----------
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Optimus Protocol API"))
	})

	r.Mount("/did", didCtrl.Routes())
	r.Mount("/bnpl", bnplCtrl.Routes())
	r.Mount("/dao", daoCtrl.Routes())
	r.Mount("/loan", loanCtrl.Routes())
	r.Mount("/vault", tvCtrl.Routes())

	// ---------- server ----------
	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Optimus protocol server listening on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}