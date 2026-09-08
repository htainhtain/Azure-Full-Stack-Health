package main

import (
	"log"
	"net/http"
	"os"
	backendHealth "server/backendhealth"
	"server/keyvault"
	"server/middleware"
	"server/redis"
	"server/sqlhealth"
	"time"

	_ "github.com/microsoft/go-mssqldb"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health/ping", backendHealth.BackendHealth)
	mux.HandleFunc("/health/sql", sqlhealth.SqlHealth)
	mux.HandleFunc("/health/redis", redis.RedisHealth)
	mux.HandleFunc("/health/keyvault", keyvault.KeyVaultHealth)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      middleware.LoggingMiddleware(mux),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("Backend running on port %s", port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
