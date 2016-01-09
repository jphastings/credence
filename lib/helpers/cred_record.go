package helpers

import (
  "github.com/golang/protobuf/proto"
  "github.com/jphastings/credence/lib/models"
  "github.com/jphastings/credence/lib/definitions/credence"
)

func StoreCredUnknownAuthor(cred *credence.Cred) {
  // TODO: Find author
  panic("Not implemented")
  // sql.NullInt64{Valid: false} may be useful?
}

func StoreCredWithAuthor(cred *credence.Cred, author models.User) {
  credBytes, _ := proto.Marshal(cred)

  var keys []models.CredKey
  for _, key := range cred.Keys {
    keyObj := models.CredKey { Key: key }
    keys = append(keys, keyObj)
  }

  credRecord := models.CredRecord{
    Author: author,
    CredBytes: credBytes,
    StatementHash: StatementHash(cred),
    Keys: keys,

    NoComment: cred.Assertion == credence.Cred_NO_COMMENT,
    IsTrue: cred.Assertion == credence.Cred_IS_TRUE,
    IsFalse: cred.Assertion == credence.Cred_IS_FALSE,
    IsAmbiguous: cred.Assertion == credence.Cred_IS_AMBIGUOUS,
  }

  db := models.DB()
  db.Create(&credRecord)
}
