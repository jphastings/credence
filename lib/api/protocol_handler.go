package api

import (
  "fmt"
  "net/http"
)

func ProtocolHandlerHandler(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case "GET": ProtocolHandler(w, r)
    default: w.WriteHeader(http.StatusMethodNotAllowed)
  }
}

func ProtocolHandler(w http.ResponseWriter, r *http.Request) {
  path := r.URL.Query()["uri"][0]
  if path[0:10] != "creds/info" {
    path = ""
  }
  url := fmt.Sprintf("/%s", path)
  w.Header().Set("Location", url)
  w.WriteHeader(http.StatusMovedPermanently)
}