package api

import (
  "sync"
  "net/http"
)

func StartAPI(wg sync.WaitGroup) {
  defer wg.Done()

  http.HandleFunc("/creds", CredHandler)
  panic(http.ListenAndServe(":8808", nil))
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusMethodNotAllowed)
}