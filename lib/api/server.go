package api

import (
  "fmt"
  "log"
  "sync"
  "net/http"
  "github.com/toqueteos/webbrowser"
  "github.com/jphastings/credence/lib/config"
  "github.com/jphastings/credence/lib/api/creds"
)

func StartAPI(wg sync.WaitGroup) {
  defer wg.Done()

  config := config.Read()
  listenUri := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)

  mux := http.NewServeMux()

  mux.HandleFunc("/creds/info", api.RawCredHandler)
  mux.HandleFunc("/creds/info/", api.InfoCredHandler)
  mux.HandleFunc("/creds", CredHandler)
  mux.HandleFunc("/connect", ConnectHandler)
  mux.HandleFunc("/ping", PingHandler)
  mux.HandleFunc("/protocol-handler", ProtocolHandlerHandler)

  static := http.FileServer(http.Dir("htdocs"))
  mux.Handle("/", static)

  server := &http.Server{
    Addr: listenUri,
    Handler: mux,
  }
  
  log.Println("Webservice will start on", listenUri)
  if config.Application.OpenWebUIOnStart {
    log.Println("Opening web browser…")
    // TODO: Create proper welcome page
    webbrowser.Open(fmt.Sprintf("%s",listenUri))
  }
  panic(server.ListenAndServe())
}
