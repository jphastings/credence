package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/mattn/go-sqlite3"
)

var db gorm.DB

func init() {
  var err error
  db, err = gorm.Open("sqlite3", "credence.db")

  if err != nil {
    panic(err)
  }

  db.AutoMigrate(
    &User{},
    &CredRecord{},
    &CredKey{},
  )
}

func DB() {
  db.DB()
}