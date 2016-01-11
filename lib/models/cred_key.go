package models

import (
  "github.com/jphastings/credence/lib/config"
  "github.com/jphastings/credence/lib/definitions/credence"
)

type CredKey struct {
  CredRecord CredRecord
  CredRecordID int `sql:"index"`
  Key string
}

// TODO: multiple LIKEs at once

func SearchCredKeys(key string) []*credence.Cred {
  var (
    results []*credence.Cred
    rows []*CredRecord
  )

  db.
    Select("cred_bytes").
    Joins("left join cred_keys on cred_keys.cred_record_id = cred_records.id").
    Where("key LIKE ?", key).
    Find(&rows)

  for _, credRecord := range rows {
    results = append(results, credRecord.Cred())
  }

  return results
}

func SearchCredKeysBreakdown(key string) []*credence.SearchResult_KeyBreakdown {
  var results []*credence.SearchResult_KeyBreakdown

  cfg := config.Read()

  var dbSpecificSelect string
  switch cfg.DB.Type {
  case "postgres":
    dbSpecificSelect = "k.key, r.statement_hash, sum(r.no_comment::int * coalesce(a.weight, 1)), sum(r.is_true::int * coalesce(a.weight, 1)), sum(r.is_false::int * coalesce(a.weight, 1)), sum(r.is_ambiguous::int * coalesce(a.weight, 1)), sum(coalesce(a.weight, 0))"
  case "sqlite3":
    dbSpecificSelect = "k.key, r.statement_hash, sum(r.no_comment * coalesce(a.weight, 1)), sum(r.is_true * coalesce(a.weight, 1)), sum(r.is_false * coalesce(a.weight, 1)), sum(r.is_ambiguous * coalesce(a.weight, 1)), sum(coalesce(a.weight, 0))"
  }

  rows, err := db.
    Table("cred_records r").
    Select(dbSpecificSelect).
    Joins("left join cred_keys k on k.cred_record_id = r.id left join users a on r.author_id = a.id").
    Where("k.key LIKE ?", key).
    Group("k.key, r.statement_hash").
    Rows()

  if err != nil {
    panic(err)
  }

  for rows.Next() {
    var (
      key string
      breakdown credence.SearchResult_AssertionBreakdown
    )
    rows.Scan(
      &key,
      &breakdown.StatementHash,
      &breakdown.NoComment,
      &breakdown.IsTrue,
      &breakdown.IsFalse,
      &breakdown.IsAmbiguous,
      &breakdown.Recognised,
    )

    assertions := []*credence.SearchResult_AssertionBreakdown{&breakdown}
    keyBreakdown := credence.SearchResult_KeyBreakdown{
      Key: key,
      Assertions: assertions,
    }
    results = append(results, &keyBreakdown)
  }

  return results
}