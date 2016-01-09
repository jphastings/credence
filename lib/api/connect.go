package api

import (
  "net/http"
  "github.com/jphastings/credence/lib/receive"
)

func ConnectHandler(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case "POST":
      uris := r.URL.Query()["uri"]

      for _, uri := range uris {
        receive.ConnectToBroadcaster(uri)
      }

      w.WriteHeader(http.StatusOK)
    default: MethodNotAllowed(w, r)
  }
}