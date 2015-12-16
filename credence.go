package main

import (
  "sync"
  "github.com/jphastings/credence/lib/api"
  "github.com/jphastings/credence/lib/broadcast"
)

func main() {
  var wg sync.WaitGroup
  wg.Add(2)

  go broadcast.StartBroadcaster(wg)
  go api.StartAPI(wg)

  wg.Wait()
}