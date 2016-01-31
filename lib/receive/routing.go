package receive

import (
  "github.com/golang/protobuf/proto"
  "github.com/jphastings/credence/lib/config"
  "github.com/jphastings/credence/lib/models"
  "github.com/jphastings/credence/lib/helpers"
  "github.com/jphastings/credence/lib/definitions/credence"
)

func RouteMessage(message *credence.Message) {
  rebroadcast := false

  cred := message.GetCred()
  if cred != nil {
    rebroadcast = ProcessInboundCred(cred)
  }

  searchRequest := message.GetSearchRequest()
  if searchRequest != nil {
    rebroadcast = ProcessInboundSearchRequest(searchRequest)
  }

  if rebroadcast {
    BroadcastMessage(message)
  }
}

func ProcessInboundCred(cred *credence.Cred) bool {
  notSeenBefore, _ := helpers.StoreCredUnknownAuthor(cred)
  return notSeenBefore
}

// Return: whether or not to rebroadcast the message
func ProcessInboundSearchRequest(searchRequest *credence.SearchRequest) bool {
  config := config.Read()
  // TODO: Store search request

  if searchRequest.Proximity <= config.SearchRequests.ForwardProximityLimit {
    searchRequest.Proximity += 1

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

    return true
  }
  return false
}

func BroadcastMessage(message *credence.Message) {
  msgBytes, _ := proto.Marshal(message)
  _, err := broadcaster.Write(msgBytes)
  if err != nil {
    panic(err)
  }
}