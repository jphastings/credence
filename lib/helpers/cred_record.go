package helpers

import (
  "time"
  "github.com/golang/protobuf/proto"
  "github.com/jphastings/credence/lib/models"
  "github.com/jphastings/credence/lib/definitions/credence"
)

func StoreCredUnknownAuthor(cred *credence.Cred) bool {
  author := DetectAuthor(cred)
  return StoreCredWithAuthor(cred, author)
}

func StoreCredWithAuthor(cred *credence.Cred, author models.User) bool {
  credBytes, _ := proto.Marshal(cred)

  db := models.DB()
  var credRecord models.CredRecord
  db.FirstOrInit(&credRecord, models.CredRecord{CredBytes: credBytes})

  if db.NewRecord(credRecord) {
    var keys []models.CredKey
    for _, key := range cred.Keys {
      keyObj := models.CredKey { Key: key }
      keys = append(keys, keyObj)
    }

    credRecord.Author = author
    credRecord.StatementHash = StatementHash(cred)
    credRecord.Keys = keys
    credRecord.ReceivedAt = time.Now()
    
    credRecord.NoComment = cred.Assertion == credence.Cred_NO_COMMENT
    credRecord.IsTrue = cred.Assertion == credence.Cred_IS_TRUE
    credRecord.IsFalse = cred.Assertion == credence.Cred_IS_FALSE
    credRecord.IsAmbiguous = cred.Assertion == credence.Cred_IS_AMBIGUOUS

    db.Save(&credRecord)
    return true
  } else {
    return false
  }
}
