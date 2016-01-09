package api

import (
  "net/http"
  "github.com/jphastings/credence/lib/api/creds"
)

func CredHandler(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case "GET": api.SearchCredHandler(w, r)
    case "POST": api.CreateCredHandler(w, r)
    default: MethodNotAllowed(w, r)
  }
}
