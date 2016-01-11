package broadcast

import (
  "fmt"
  "log"
  "sync"
  "github.com/zeromq/goczmq"
  "github.com/jphastings/credence/lib/config"
)

var receiver *goczmq.Sock
var broadcaster *goczmq.Sock

func Setup() {
  config := config.Read()
  broadcastUri := fmt.Sprintf("tcp://%s:%d", config.Broadcaster.Host, config.Broadcaster.Port)

  var err error
  broadcaster, err = goczmq.NewPub(broadcastUri)
  if err != nil {
    panic(err)
  }
  fmt.Println("Broadcaster started on", broadcastUri)

  receiver, err = goczmq.NewPull("inproc://broadcast")
  if err != nil {
    panic(err)
  }
}

func StartBroadcaster(wg sync.WaitGroup) {
  defer wg.Done()

  msgBytes := make([]byte, 524288) // 0.5 Mb

  for {
    _, err := receiver.Read(msgBytes)
    if err != nil {
      panic(err)
      continue
    }

    log.Print("Broadcasting message")

    _, err = broadcaster.Write(msgBytes)
    if err != nil {
      panic(err)
    }
  }
}