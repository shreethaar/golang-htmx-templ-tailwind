package main
import (
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/joho/godotenv"
    "log/slog"
    "log"
    "os"
)

func main() {
    if err:=godotenv.Load(); err!=nil {
        log.Fatal(err)
    }

    router:=chi.NewMux()
    router.Get("/foo",handleFoo)
    listenAddr:=os.Getenv("LISTEN_ADDR")
    slog.Info("HTTP server started","listenAddr",listenAddr)
    http.ListenAndServe(listenAddr,router)

}

func handleFoo(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("foo"))

}
