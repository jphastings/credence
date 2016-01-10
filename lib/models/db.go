package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/mattn/go-sqlite3"
  _ "github.com/lib/pq"
  "github.com/jphastings/credence/lib/config"
)

var db gorm.DB

func Setup() {
  var err error
  cfg := config.Read()

  connectionString := cfg.DB.ConnectionString
  if cfg.DB.Type == "sqlite3" {
    connectionString = config.ConfigFile(connectionString)
  }
  db, err = gorm.Open(cfg.DB.Type, connectionString)

  if err != nil {
    panic(err)
  }

  db.AutoMigrate(
    &User{},
    &CredRecord{},
    &CredKey{},
  )

  db.Model(&CredRecord{}).AddUniqueIndex("idx_cred_bytes", "cred_bytes")
}

func DB() gorm.DB {
  return db
}
