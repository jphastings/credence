package web

import (
  "fmt"
  "log"
  "sync"
  "net/http"
  "github.com/toqueteos/webbrowser"
  "github.com/jphastings/credence/lib/config"
  "github.com/jphastings/credence/lib/web/creds"
)

func StartWeb(wg sync.WaitGroup) {
  defer wg.Done()

  cfg := config.Read()
  listenUri := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)

  mux := http.NewServeMux()

  mux.HandleFunc("/creds/info", web.RawCredHandler)
  mux.HandleFunc("/creds/info/", web.InfoCredHandler)
  mux.HandleFunc("/creds", CredHandler)

  mux.HandleFunc("/users/", UsersHandler)

  mux.HandleFunc("/connect", ConnectHandler)
  mux.HandleFunc("/ping", PingHandler)
  mux.HandleFunc("/protocol-handler", ProtocolHandlerHandler)

  static := http.FileServer(http.Dir(config.AssetDir("htdocs")))
  mux.Handle("/", static)

  server := &http.Server{
    Addr: listenUri,
    Handler: mux,
  }
  
  log.Println("Webservice will start on", listenUri)
  if cfg.Application.OpenWebUIOnStart {
    log.Println("Opening web browser…")
    // TODO: Create proper welcome page
    webbrowser.Open(fmt.Sprintf("%s",listenUri))
  }
  panic(server.ListenAndServe())
}
