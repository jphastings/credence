package web

import (
  "net/http"
  "github.com/jphastings/credence/lib/helpers"
)

func InfoCredHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != "GET" {
    w.WriteHeader(http.StatusMethodNotAllowed)
    return
  }

  // TODO: This is a hack, need proper routing
  credHash := r.URL.Path[12:]

  credRecord, found := helpers.CredRecordFromCredHash(credHash)
  if !found {
    w.WriteHeader(http.StatusNotFound)
    // TODO: Pretty 404
    return
  }

  w.WriteHeader(http.StatusOK)
  helpers.ModelNegotiator().Negotiate(w, r, &credRecord)
}