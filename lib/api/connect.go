package api

import (
  "net/http"
  "github.com/jphastings/credence/lib/models"
  "github.com/jphastings/credence/lib/receive"
)

func ConnectHandler(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case "POST":
      uris := r.URL.Query()["uri"]

      for _, uri := range uris {
        err := receive.ConnectToBroadcaster(uri)
        if err == nil {
          db := models.DB()
          db.FirstOrCreate(new(models.Peer), models.Peer{Server: uri})
        } else {
          w.WriteHeader(http.StatusBadRequest)
          return
        }
      }

      w.WriteHeader(http.StatusOK)
    default: MethodNotAllowed(w, r)
  }
}