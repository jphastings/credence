package web

import (
  "net/http"
  "github.com/jphastings/credence/lib/helpers"
  "github.com/jphastings/credence/lib/web/creds"
)

func CredHandler(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case "GET": web.SearchCredHandler(w, r)
    case "POST":
      if helpers.HasPrivateKey() {
        web.CreateCredHandler(w, r)
      } else {
        w.WriteHeader(http.StatusMethodNotAllowed)
      }
    default: w.WriteHeader(http.StatusMethodNotAllowed)
  }
}
