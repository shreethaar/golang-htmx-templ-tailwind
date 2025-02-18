package handlers

import (
    "net/http"
    "fmt"
)
func HandleFoo(w http.ResponseWriter,r *http.Request) error {
    w.Write([]byte("foo"))
    return fmt.Errorf("helpme")

}
