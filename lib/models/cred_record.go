package models

import (
  "time"
  "github.com/golang/protobuf/proto"
  "github.com/jphastings/credence/lib/config"
  "github.com/jphastings/credence/lib/definitions/credence"
)

type CredRecord struct {
  ID uint `gorm:"primary_key"`
  SourceUri string
  CredBytes []byte 
  // The detected author of this cred
  Author User
  AuthorID uint
  // An identifier for this statement in this context; the hash of the source uri and statement
  StatementHash string
  // An identifier for the specific Cred
  CredHash string `sql:unique_index`
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


// TODO: multiple LIKEs at once
func SearchCreds(key string) []*credence.Cred {
  var (
    results []*credence.Cred
    rows []*CredRecord
  )

  db.
    Select("cred_bytes").
    Where("source_uri LIKE ?", key).
    Find(&rows)

  for _, credRecord := range rows {
    results = append(results, credRecord.Cred())
  }

  return results
}

func SearchCredsBreakdown(key string) []*credence.SearchResult_SourceBreakdown {
  var results []*credence.SearchResult_SourceBreakdown

  cfg := config.Read()

  var dbSpecificSelect string
  switch cfg.DB.Type {
  case "postgres":
    dbSpecificSelect = "source_uri, statement_hash, sum(no_comment::int * coalesce(a.weight, 1)), sum(is_true::int * coalesce(a.weight, 1)), sum(is_false::int * coalesce(a.weight, 1)), sum(is_ambiguous::int * coalesce(a.weight, 1)), sum(coalesce(a.weight, 0))"
  case "sqlite3":
    dbSpecificSelect = "source_uri, statement_hash, sum(no_comment * coalesce(a.weight, 1)), sum(is_true * coalesce(a.weight, 1)), sum(is_false * coalesce(a.weight, 1)), sum(is_ambiguous * coalesce(a.weight, 1)), sum(coalesce(a.weight, 0))"
  }

  rows, err := db.
    Table("cred_records").
    Select(dbSpecificSelect).
    Joins("left join users a on author_id = a.id").
    Where("source_uri LIKE ?", key).
    Group("source_uri, statement_hash").
    Rows()

  if err != nil {
    panic(err)
  }

  for rows.Next() {
    var (
      sourceUri string
      breakdown credence.SearchResult_AssertionBreakdown
    )
    rows.Scan(
      &sourceUri,
      &breakdown.StatementHash,
      &breakdown.NoComment,
      &breakdown.IsTrue,
      &breakdown.IsFalse,
      &breakdown.IsAmbiguous,
      &breakdown.Recognised,
    )

    assertions := []*credence.SearchResult_AssertionBreakdown{&breakdown}
    sourceBreakdown := credence.SearchResult_SourceBreakdown{
      SourceUri: sourceUri,
      Assertions: assertions,
    }
    results = append(results, &sourceBreakdown)
  }

  return results
}
