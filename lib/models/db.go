package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/mattn/go-sqlite3"
  "github.com/jphastings/credence/lib/config"
)

var db gorm.DB

func init() {
  var err error
  dbPath := config.ConfigFile("credence.db")
  db, err = gorm.Open("sqlite3", dbPath)

  if err != nil {
    panic(err)
  }

  db.AutoMigrate(
    &User{},
    &CredRecord{},
    &CredKey{},
  )
}

func DB() gorm.DB {
  return db
}
