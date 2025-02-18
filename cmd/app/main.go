package main
import (
    "net/http"
    "github.com/go-chi/chi/v5"
    "log/slog"
)

func main() {
    router:=chi.NewMux()
    router.Get("/foo",handleFoo)
    listenAddr:=":3000"
    slog.Info("HTTP server started","listenAddr",listenAddr)
    http.ListenAndServe(listenAddr,router)

}

func handleFoo(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("foo"))

}
