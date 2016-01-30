package api

import (
  "io"
  "time"
  "net/http"
  "github.com/zeromq/goczmq"
  "github.com/golang/protobuf/proto"
  "github.com/golang/protobuf/jsonpb"
  "github.com/jphastings/credence/lib/models"
  "github.com/jphastings/credence/lib/helpers"
  "github.com/jphastings/credence/lib/definitions/credence"
)

func SearchCredHandler(w http.ResponseWriter, r *http.Request) {
  // TODO: Set timeout header for after search result comes back
  
  broadcastSearch := r.URL.Query()["broadcast"] != nil
  queryKeys := r.URL.Query()["key"]

  if len(queryKeys) == 0 {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  searchRequest := &credence.SearchRequest{
    Keys: queryKeys,
    Timestamp: time.Now().Unix(),
  }

  if broadcastSearch {
    BroadcastSearch(searchRequest)
  }

  searchResult := &credence.SearchResult{}

  // TODO: Do DB query for all keys at once
  for _, key := range searchRequest.Keys {
    for _, keyBreakdown := range models.SearchCredsBreakdown(key) {
      searchResult.Results = append(searchResult.Results, keyBreakdown)
    }
  }

  helpers.DeduplicateKeys(searchResult)
  
  // Output
  marshaler := jsonpb.Marshaler{}
  json, _ := marshaler.MarshalToString(searchResult)

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  io.WriteString(w, json)
}

func BroadcastSearch(searchRequest *credence.SearchRequest) {
  // Set up the broadcaster
  broadcaster, err := goczmq.NewPush("inproc://broadcast")
  if err != nil {
      panic(err)
  }
  defer broadcaster.Destroy()

  // Create the broadcastable bytes
  msg := &credence.Message{
    Type: &credence.Message_SearchRequest{
      SearchRequest: searchRequest,
    },
  }

  msgBytes, err := proto.Marshal(msg)
  if err != nil {
    panic(err)
  }

  _, err = broadcaster.Write(msgBytes)
  if err != nil {
      panic(err)
  }
}