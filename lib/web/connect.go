package web

import (
  "net/http"
  "github.com/jphastings/credence/lib/models"
  "github.com/jphastings/credence/lib/receive"
)

// TODO: Allow addition of broadcatchers too
func ConnectHandler(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case "POST":
      uris := r.URL.Query()["uri"]

      for _, uri := range uris {
        err := receive.ConnectToBroadcaster(uri)
        if err == nil {
          db := models.DB()
          db.FirstOrCreate(new(models.Peer), models.Peer{Server: uri, IsBroadcatcher: false})
        } else {
          w.WriteHeader(http.StatusBadRequest)
          return
        }
      }

      w.WriteHeader(http.StatusOK)
    default: w.WriteHeader(http.StatusMethodNotAllowed)
  }
}