package receive

import (
  "fmt"
  "log"
  "sync"
  "github.com/zeromq/goczmq"
  "github.com/golang/protobuf/proto"
  "github.com/jphastings/credence/lib/config"
  "github.com/jphastings/credence/lib/models"
  "github.com/jphastings/credence/lib/helpers"
  "github.com/jphastings/credence/lib/definitions/credence"
)

var receiver *goczmq.Sock
var broadcaster *goczmq.Sock

func Setup() {
  var err error
  broadcaster, err = goczmq.NewPush("inproc://broadcast")
  if err != nil {
    panic(err)
  }

  receiver = goczmq.NewSock(goczmq.Sub)
  receiver.SetSubscribe("")
  if err != nil {
    panic(err)
  }
}

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

func ConnectToBroadcaster(uri string) {
  broadcasterUri := fmt.Sprintf("tcp://%s", uri)
  log.Println("Connecting to", broadcasterUri)
  err := receiver.Connect(broadcasterUri)
  if err != nil {
    log.Print(err)
  }
}

func RouteMessage(message *credence.Message) {
  config := config.Read()

  log.Println("Routing message", message)

  cred := message.GetCred()
  if cred != nil {
    if helpers.StoreCredUnknownAuthor(cred) {
      BroadcastMessage(message)
    }
  }

  searchRequest := message.GetSearchRequest()
  if searchRequest != nil {
    // TODO: Store search request

    if searchRequest.Proximity <= config.SearchRequests.ForwardProximityLimit {
      searchRequest.Proximity += 1
      BroadcastMessage(message)

      for _, key := range searchRequest.Keys {
        for _, cred := range models.SearchCredKeys(key) {
          credMsg := &credence.Message{
            Type: &credence.Message_Cred{
              Cred: cred,
            },
          }

          BroadcastMessage(credMsg)
        }
      }
    }
  }
}

func BroadcastMessage(message *credence.Message) {
  msgBytes, _ := proto.Marshal(message)
  _, err := broadcaster.Write(msgBytes)
  if err != nil {
    panic(err)
  }
}