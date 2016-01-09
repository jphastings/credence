package models

import (
  "time"
  "database/sql"
  "github.com/golang/protobuf/proto"
  "github.com/jphastings/credence/lib/definitions/credence"
)

type CredRecord struct {
  ID int
  Keys   []CredKey
  CredBytes []byte 
  // The detected author of this cred
  Author User
  AuthorID sql.NullInt64
  // An identifier for this statement in this context; the hash of the keys and statement
  StatementHash string
  // The time this cred was received
  ReceivedAt time.Time

  NoComment bool
  IsTrue bool
  IsFalse bool
  IsAmbiguous bool

  Seen int `sql:"DEFAULT:1"`
}

func (credRecord CredRecord) Cred() (*credence.Cred) {
 cred := &credence.Cred{}
 err := proto.Unmarshal(credRecord.CredBytes, cred)
 if err != nil {
  panic(err)
 }
 return cred
}
