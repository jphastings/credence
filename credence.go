package main

import (
  "sync"
  "github.com/jphastings/credence/lib/config"
  "github.com/jphastings/credence/lib/models"
  "github.com/jphastings/credence/lib/api"
  "github.com/jphastings/credence/lib/broadcast"
)

func main() {
  config.Setup()
  config.Read()
  models.Setup()

  var wg sync.WaitGroup
  wg.Add(2)

  go broadcast.StartBroadcaster(wg)
  go api.StartAPI(wg)

  wg.Wait()
}