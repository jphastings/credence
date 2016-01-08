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
}

func (credRecord CredRecord) Cred() (*credence.Cred) {
 cred := &credence.Cred{}
 err := proto.Unmarshal(credRecord.CredBytes, cred)
 if err != nil {
  panic(err)
 }
 return cred
}

func StoreCredUnknownAuthor(cred *credence.Cred) {
  // TODO: Find author
  panic("Not implemented")
  // sql.NullInt64{Valid: false} may be useful?
}

func StoreCredWithAuthor(cred *credence.Cred, author User) {
  credBytes, _ := proto.Marshal(cred)

  var keys []CredKey
  for _, key := range cred.Keys {
    keyObj := CredKey { Key: key }
    keys = append(keys, keyObj)
  }

  credRecord := CredRecord{
    Author: author,
    CredBytes: credBytes,
    StatementHash: cred.StatementHash(),
    Keys: keys,

    NoComment: cred.Assertion == credence.Cred_NO_COMMENT,
    IsTrue: cred.Assertion == credence.Cred_IS_TRUE,
    IsFalse: cred.Assertion == credence.Cred_IS_FALSE,
    IsAmbiguous: cred.Assertion == credence.Cred_IS_AMBIGUOUS,
  }

  db.Create(&credRecord)
}
