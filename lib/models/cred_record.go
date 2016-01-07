package models

import (
  "time"
  "crypto/sha1"
  "encoding/hex"
  "github.com/golang/protobuf/proto"
  "github.com/jphastings/credence/lib/definitions/credence"
)

type CredRecord struct {
  ID int
  Keys   []CredKey
  CredBytes []byte 
  // The detected author of this cred
  Author User
  // An identifier for this statement in this context; the hash of the keys and statement
  Mark string
  // The time this cred was received
  ReceivedAt time.Time

  NoComment bool
  IsTrue bool
  IsFalse bool
  IsAmbiguous bool
}

func (credRecord CredRecord) Cred() (*credence.Cred) {
 cred := &credence.Cred{}
 err := proto.Unmarshal(credRecord.CredBytes, cred)
 if err != nil {
  panic(err)
 }
 return cred
}

func StoreCred (cred *credence.Cred) {
  // TODO: Determine Author
  credBytes, _ := proto.Marshal(cred)

  var keys []CredKey
  for _, key := range cred.Keys {
    keyObj := CredKey { Key: key }
    keys = append(keys, keyObj)
  }

  credRecord := CredRecord{
    CredBytes: credBytes,
    Mark: MarkFor(cred),
    Keys: keys,

    NoComment: cred.Assertion == credence.Cred_NO_COMMENT,
    IsTrue: cred.Assertion == credence.Cred_IS_TRUE,
    IsFalse: cred.Assertion == credence.Cred_IS_FALSE,
    IsAmbiguous: cred.Assertion == credence.Cred_IS_AMBIGUOUS,
  }

  db.Create(&credRecord)
}

func MarkFor(cred *credence.Cred) string {
  markCred := &credence.Cred{
    Keys: cred.Keys,
    Statement: cred.Statement,
  }

  markCredBytes, err := proto.Marshal(markCred)
  if err != nil {
    // TODO: Deal with error
    panic(err)
  }

  hash := sha1.New()
  hash.Write(markCredBytes) 
  return hex.EncodeToString(hash.Sum(nil))
}