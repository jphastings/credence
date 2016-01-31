package receive

import (
  "github.com/golang/protobuf/proto"
  "github.com/jphastings/credence/lib/config"
  "github.com/jphastings/credence/lib/models"
  "github.com/jphastings/credence/lib/helpers"
  "github.com/jphastings/credence/lib/definitions/credence"
)

func RouteMessage(message *credence.Message) {
  config := config.Read()

  cred := message.GetCred()
  if cred != nil {
    newCred, _ := helpers.StoreCredUnknownAuthor(cred)
    if newCred {
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
        for _, cred := range models.SearchCreds(key) {
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