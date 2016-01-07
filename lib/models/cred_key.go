package models

import (
  "github.com/jphastings/credence/lib/definitions/credence"
)

type CredKey struct {
  CredRecord CredRecord
  CredRecordID int `sql:"index"`
  Key string
}

func SearchCredKeys(key string) []*credence.SearchResult_KeyBreakdown {
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

    // TODO: Don't duplicate key breakdowns for the same key
    assertions := []*credence.SearchResult_AssertionBreakdown{&breakdown}
    keyBreakdown := credence.SearchResult_KeyBreakdown{
      Key: key,
      Assertions: assertions,
    }
    results = append(results, &keyBreakdown)
  }

  return results
}