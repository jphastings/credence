package receive

import (
  "log"
  "sync"
  "github.com/golang/protobuf/proto"
  "github.com/jphastings/credence/lib/definitions/credence"
)

func StartReceiver(wg sync.WaitGroup) {
  defer wg.Done()

  msgBytes := make([]byte, 524288) // 0.5 Mb

  for {
    _, err := receiver.Read(msgBytes)
    if err != nil {
      log.Print(err)
      continue
    }
    log.Print("Message received")

    message := &credence.Message{}
    err = proto.Unmarshal(msgBytes, message)
    if err != nil {
      // TODO: Why is this failing?
      log.Print(err)
    }
    RouteMessage(message)
  }
}
