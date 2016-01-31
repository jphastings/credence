package main

import (
  "sync"
  "github.com/jphastings/credence/lib/config"
  "github.com/jphastings/credence/lib/models"
  "github.com/jphastings/credence/lib/web"
  "github.com/jphastings/credence/lib/helpers"
  "github.com/jphastings/credence/lib/broadcast"
  "github.com/jphastings/credence/lib/receive"
)

func main() {
  config.Setup()
  models.Setup()
  broadcast.Setup()
  receive.Setup()
  helpers.SavePublicKeyIfNeccessary()

  var wg sync.WaitGroup
  wg.Add(4)

  go broadcast.StartBroadcaster(wg)
  go receive.StartReceiver(wg)
  go receive.StartBroadcatcher(wg)
  go web.StartWeb(wg)

  wg.Wait()
}