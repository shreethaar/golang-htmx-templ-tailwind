package handlers

import (
    "net/http"
 //   "fmt"
    "trial-run/views/foo"
)

func HandleFoo(w http.ResponseWriter,r *http.Request) error {
    return foo.Index().Render(r.Context(),w)
    //return fmt.Errorf("helpme")

}
