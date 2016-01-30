package helpers

import (
  "time"
  "github.com/golang/protobuf/proto"
  "github.com/jphastings/credence/lib/models"
  "github.com/jphastings/credence/lib/definitions/credence"
)

func StoreCredUnknownAuthor(cred *credence.Cred) bool {
  author, _ := DetectAuthor(cred)
  // TODO: propagate error 
  return StoreCredWithAuthor(cred, author)
}

func StoreCredWithAuthor(cred *credence.Cred, author models.User) bool {
  credBytes, _ := proto.Marshal(cred)

  db := models.DB()
  var credRecord models.CredRecord
  db.FirstOrInit(&credRecord, models.CredRecord{CredBytes: credBytes})

  if db.NewRecord(credRecord) {
    credRecord.Author = author
    credRecord.StatementHash = StatementHash(cred)
    credRecord.SourceUri = cred.SourceUri
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
