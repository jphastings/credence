package models

import (
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

  rows, err := db.
    Model(CredRecord{}).
    Select("key, statement_hash, sum(no_comment), sum(is_true), sum(is_false), sum(is_ambiguous), count(author_id)").
    Joins("left join cred_keys on cred_keys.cred_record_id = cred_records.id").
    Where("key LIKE ?", key).
    Group("key, statement_hash").
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