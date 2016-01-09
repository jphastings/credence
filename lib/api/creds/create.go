package api

import (
  "io"
  "log"
  "time"
  "net/http"
  "github.com/zeromq/goczmq"
  "github.com/golang/protobuf/proto"
  "github.com/golang/protobuf/jsonpb"
  "github.com/jphastings/credence/lib/models"
  "github.com/jphastings/credence/lib/helpers"
  "github.com/jphastings/credence/lib/definitions/credence"
)

func CreateCredHandler(w http.ResponseWriter, r *http.Request) {
  cred := &credence.Cred{}
  if err := jsonpb.Unmarshal(r.Body, cred); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  // Set attributes
  cred.Timestamp = time.Now().Unix()
  err := helpers.SetSignature(cred)
  if err != nil {
    panic(err)
  }

  // Store in the DB
  helpers.StoreCredWithAuthor(cred, models.Me())

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