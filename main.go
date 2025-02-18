package main

import (
    "net/http"
    "trial-run/handlers"  
    "github.com/go-chi/chi/v5"
    "github.com/joho/godotenv"
    "log/slog"
    "log"
    "os"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatal(err)
    }
    
    router := chi.NewMux()
    router.Get("/fooOOO", handlers.Make(handlers.HandleFoo))
    
    listenAddr := os.Getenv("LISTEN_ADDR")
    slog.Info("HTTP server started", "listenAddr", listenAddr)
    http.ListenAndServe(listenAddr, router)
}
