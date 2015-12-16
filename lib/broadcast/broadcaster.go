package broadcast

import (
  "log"
  "sync"
  "github.com/zeromq/goczmq"
)

func StartBroadcaster(wg sync.WaitGroup) {
  defer wg.Done()

  // Set up the broadcaster
  broadcaster, err := goczmq.NewPush("tcp://0.0.0.0:9099")
  if err != nil {
    panic(err)
  }
  defer broadcaster.Destroy()

  receiver, err := goczmq.NewPull("inproc://broadcast")
  if err != nil {
    panic(err)
  }
  defer receiver.Destroy()
  log.Println("Broadcast pull created and connected")

  msgBytes := make([]byte, 524288) // 0.5 Mb

  for {
    // TODO: Do I need to reset buf?
    _, err := receiver.Read(msgBytes)
    if err != nil {
      log.Print(err)
      continue
    }

    _, err = broadcaster.Write(msgBytes)
    if err != nil {
      panic(err)
    }
  }
}