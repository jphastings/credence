package api

import (
  "io"
  "log"
  "time"
  "net/http"
  "github.com/zeromq/goczmq"
  "github.com/golang/protobuf/proto"
  "github.com/golang/protobuf/jsonpb"
  "github.com/jphastings/credence/lib/definitions/credence"
  "github.com/jphastings/credence/lib/models"
)

var broadcaster *goczmq.Channeler

func init() {
  // TODO: Pull out the broadcaster
}

func CredHandler(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case "GET": SearchCredHandler(w, r)
    case "POST": CreateCredHandler(w, r)
    default: MethodNotAllowed(w, r)
  }
}

func SearchCredHandler(w http.ResponseWriter, r *http.Request) {
  // TODO: Set timeout header for after search result comes back
  
  queryKeys := r.URL.Query()["key"]

  if len(queryKeys) == 0 {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  searchRequest := &credence.SearchRequest{
    Keys: queryKeys,
    Timestamp: time.Now().Unix(),
  }

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

  searchResult := &credence.SearchResult{}

  for _, key := range searchRequest.Keys {
    for _, keyBreakdown := range models.SearchCredKeys(key) {
      searchResult.Results = append(searchResult.Results, keyBreakdown)
    }
  }

  searchResult.DeduplicateKeys()
  
  // Output
  marshaler := jsonpb.Marshaler{}
  json, _ := marshaler.MarshalToString(searchResult)

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  io.WriteString(w, json)
}

func CreateCredHandler(w http.ResponseWriter, r *http.Request) {
  cred := &credence.Cred{}
  if err := jsonpb.Unmarshal(r.Body, cred); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  // Set attributes
  cred.Timestamp = time.Now().Unix()
  cred.Signature = []byte{0x1f, 0x8b}

  // Store in the DB
  models.StoreCred(cred)

  // Set up the broadcaster
  broadcaster, err := goczmq.NewPush("inproc://broadcast")
  if err != nil {
      panic(err)
  }
  defer broadcaster.Destroy()

  // Create the broadcastable bytes
  msg := &credence.Message{
    Type: &credence.Message_Cred{
      Cred: cred,
    },
  }

  msgBytes, err := proto.Marshal(msg)
  if err != nil {
    log.Print(err)
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  _, err = broadcaster.Write(msgBytes)
  if err != nil {
      panic(err)
  }

  var credMarshaled string

  // Respond over HTTP
  switch r.Header.Get("Accept") {
  case "application/vnd.google.protobuf":
    w.Header().Set("Content-Type", "application/vnd.google.protobuf")
    credBytes, _ := proto.Marshal(cred)
    credMarshaled = string(credBytes)
  case "application/json":
    w.Header().Set("Content-Type", "application/json")
    marshaler := jsonpb.Marshaler{}
    credMarshaled, _ = marshaler.MarshalToString(cred)
  default:
    w.WriteHeader(http.StatusNotAcceptable)
    return
  }
  
  w.WriteHeader(http.StatusCreated)
  io.WriteString(w, credMarshaled)
}