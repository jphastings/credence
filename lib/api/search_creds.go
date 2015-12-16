package api

import (
  "io"
  "log"
  "time"
  "net/http"
  "github.com/zeromq/goczmq"
  // "github.com/spacemonkeygo/openssl"
  "github.com/golang/protobuf/proto"
  "github.com/golang/protobuf/jsonpb"
  "github.com/jphastings/credence/lib/definitions/credence"
)

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

  searchResult := &credence.SearchResult{}

  for _, key := range queryKeys {
    keyBreakdown := &credence.SearchResult_KeyBreakdown {
      Key: key,
    }
    searchResult.Results = append(searchResult.Results, keyBreakdown)
  }
  
  // Output
  w.WriteHeader(http.StatusOK)
  marshaler := jsonpb.Marshaler{}
  json, _ := marshaler.MarshalToString(searchResult)
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

  // Respond over HTTP
  w.WriteHeader(http.StatusCreated)
  marshaler := jsonpb.Marshaler{}
  credJson, _ := marshaler.MarshalToString(cred)
  io.WriteString(w, credJson)
}