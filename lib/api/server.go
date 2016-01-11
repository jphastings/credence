package api

import (
  "fmt"
  "log"
  "sync"
  "net/http"
  "github.com/jphastings/credence/lib/config"
)

func StartAPI(wg sync.WaitGroup) {
  defer wg.Done()

  config := config.Read()
  listenUri := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)

  http.HandleFunc("/creds", CredHandler)
  http.HandleFunc("/connect", ConnectHandler)
  http.HandleFunc("/ping", PingHandler)

  static := http.FileServer(http.Dir("htdocs"))
  http.Handle("/", static)
  log.Println("Webservice started on", listenUri)
  panic(http.ListenAndServe(listenUri, nil))
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusMethodNotAllowed)
}