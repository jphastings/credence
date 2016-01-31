package helpers

import (
  "time"
  "github.com/golang/protobuf/proto"
  "github.com/jphastings/credence/lib/models"
  "github.com/jphastings/credence/lib/definitions/credence"
)

func StoreCredUnknownAuthor(cred *credence.Cred) (bool, string) {
  author, _ := DetectAuthor(cred)
  // TODO: propagate error 
  return StoreCredWithAuthor(cred, author)
}

func StoreCredWithAuthor(cred *credence.Cred, author models.User) (bool, string) {
  credBytes, _ := proto.Marshal(cred)

  db := models.DB()
  var credRecord models.CredRecord
  credHash := CredHash(cred)

  db.FirstOrInit(&credRecord, models.CredRecord{CredHash: credHash})
  newCred := db.NewRecord(credRecord)

  if newCred {
    credRecord.Author = author
    credRecord.StatementHash = StatementHash(cred)
    credRecord.CredBytes = credBytes
    credRecord.SourceUri = cred.SourceUri
    credRecord.ReceivedAt = time.Now()
    
    credRecord.NoComment = cred.Assertion == credence.Cred_NO_COMMENT
    credRecord.IsTrue = cred.Assertion == credence.Cred_IS_TRUE
    credRecord.IsFalse = cred.Assertion == credence.Cred_IS_FALSE
    credRecord.IsAmbiguous = cred.Assertion == credence.Cred_IS_AMBIGUOUS

    db.Save(&credRecord)
  }

  return newCred, credHash
}

func StatementAlreadyMade(cred *credence.Cred, user models.User) bool {
  db := models.DB()
  var credRecord models.CredRecord
  db.First(&credRecord, models.CredRecord{
    StatementHash: StatementHash(cred),
    AuthorID: user.ID,
  })
  return !db.NewRecord(credRecord)
}

func CredRecordFromCredHash(credHash string) (models.CredRecord, bool) {
  db := models.DB()
  credRecord := &models.CredRecord{}
  db.Where("cred_hash = ?", credHash).Preload("Author").First(credRecord)
  return *credRecord, !db.NewRecord(credRecord)
}
