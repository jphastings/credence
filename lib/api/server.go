package api

import (
  "fmt"
  "sync"
  "net/http"
  "github.com/jphastings/credence/lib/config"
)

func StartAPI(wg sync.WaitGroup) {
  defer wg.Done()

  config := config.Read()
  listenUri := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)

  http.HandleFunc("/creds", CredHandler)
  fmt.Println("Webservice started on", listenUri)
  panic(http.ListenAndServe(listenUri, nil))
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusMethodNotAllowed)
}