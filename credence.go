package main

import (
  "sync"
  "github.com/jphastings/credence/lib/config"
  "github.com/jphastings/credence/lib/models"
  "github.com/jphastings/credence/lib/api"
  "github.com/jphastings/credence/lib/broadcast"
  "github.com/jphastings/credence/lib/receive"
)

func main() {
  config.Setup()
  models.Setup()
  broadcast.Setup()
  receive.Setup()

  var wg sync.WaitGroup
  wg.Add(3)

  go broadcast.StartBroadcaster(wg)
  go receive.StartReceiver(wg)
  go api.StartAPI(wg)

  wg.Wait()
}