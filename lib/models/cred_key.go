package models

type CredKey struct {
  CredRecord CredRecord
  CredRecordID int `sql:"index"`
  Key string
}