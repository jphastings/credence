package api

import (
  "net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case "GET": Ping(w, r)
    default: MethodNotAllowed(w, r)
  }
}

func Ping(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.WriteHeader(http.StatusOK)
}