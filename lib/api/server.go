package api

import (
  "fmt"
  "log"
  "sync"
  "net/http"
  "github.com/jphastings/credence/lib/config"
  "github.com/jphastings/credence/lib/api/creds"
)

func StartAPI(wg sync.WaitGroup) {
  defer wg.Done()

  config := config.Read()
  listenUri := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)

  http.HandleFunc("/creds/info", api.InfoCredHandler)
  http.HandleFunc("/creds", CredHandler)
  http.HandleFunc("/connect", ConnectHandler)
  http.HandleFunc("/ping", PingHandler)
  http.HandleFunc("/protocol-handler", ProtocolHandlerHandler)

  static := http.FileServer(http.Dir("htdocs"))
  http.Handle("/", static)
  log.Println("Webservice started on", listenUri)
  panic(http.ListenAndServe(listenUri, nil))
}
