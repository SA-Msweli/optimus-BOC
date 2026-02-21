package main

import (
    "log"
    "os"
)

// Config holds all environment configuration for the Optimus backend.
// It is intentionally simple; keys are documented in protocol/README.md.

type Config struct {
    Port            string
    DatabaseURL     string
    ChainRPCURL     string
    BNPLManagerAddr string
    LoanManagerAddr string
    DAOManagerAddr  string
    DIDRegistryAddr string
    TokenVaultAddr  string
    PrivateKey      string // for signing on‑chain transactions (optional)
}

// LoadConfig reads configuration from environment variables and returns
// a Config object. It will fatally exit the process if any required
// variable is missing.
func LoadConfig() *Config {
    cfg := &Config{
        Port:            getEnv("PORT", "8000"),
        DatabaseURL:     mustGetEnv("DATABASE_URL"),
        ChainRPCURL:     mustGetEnv("CHAIN_RPC_URL"),
        BNPLManagerAddr: mustGetEnv("BNPL_MANAGER_ADDRESS"),
        LoanManagerAddr: mustGetEnv("LOAN_MANAGER_ADDRESS"),
        DAOManagerAddr:  mustGetEnv("DAO_MANAGER_ADDRESS"),
        DIDRegistryAddr: mustGetEnv("DID_REGISTRY_ADDRESS"),
        TokenVaultAddr:  mustGetEnv("TOKEN_VAULT_ADDRESS"),
        PrivateKey:      os.Getenv("PRIVATE_KEY"),
    }
    return cfg
}

func getEnv(key, defaultVal string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return defaultVal
}

func mustGetEnv(key string) string {
    val := os.Getenv(key)
    if val == "" {
        log.Fatalf("environment variable %s is required", key)
    }
    return val
}
