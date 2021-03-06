package web

import (
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
  // TODO: Prevent making the same cred a second time
  cred := &credence.Cred{}
  if err := jsonpb.Unmarshal(r.Body, cred); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  fingerprint := len(r.URL.Query()["fingerprint"]) > 0
  confirm := len(r.URL.Query()["confirm"]) > 0

  signingUser := models.Me()

  if !confirm && helpers.StatementAlreadyMade(cred, signingUser) {
    w.WriteHeader(http.StatusConflict)
    return
  }

  if fingerprint {
    cred.AuthorFingerprint = signingUser.Fingerprint
  }

  // Set attributes
  cred.Timestamp = time.Now().Unix()
  err := helpers.SetSignature(cred)
  if err != nil {
    panic(err)
  }

  // Store in the DB
  helpers.StoreCredWithAuthor(cred, signingUser)

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

  helpers.ModelNegotiator().Negotiate(w, r, cred)
}
