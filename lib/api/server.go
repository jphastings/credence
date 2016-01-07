package api

import (
  "log"
  "sync"
  "net/http"
)

func StartAPI(wg sync.WaitGroup) {
  defer wg.Done()

  http.HandleFunc("/creds", CredHandler)
  log.Print("Webservice started on 127.0.0.1:8808")
  panic(http.ListenAndServe("127.0.0.1:8808", nil))
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusMethodNotAllowed)
}