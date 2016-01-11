package api

import (
  "net/http"
  "github.com/jphastings/credence/lib/helpers"
  "github.com/jphastings/credence/lib/api/creds"
)

func CredHandler(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case "GET": api.SearchCredHandler(w, r)
    case "POST":
      if helpers.HasPrivateKey() {
        api.CreateCredHandler(w, r)
      } else {
        MethodNotAllowed(w, r)
      }
    default: MethodNotAllowed(w, r)
  }
}
