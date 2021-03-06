package models

import (
  "log"
  "time"
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
    log.Println("Using SQLite DB at", connectionString)
  }
  for {
    db, err = gorm.Open(cfg.DB.Type, connectionString)

    if err == nil {
      break
    } else {
      log.Println("Cannot connect to DB. Will try again in 2 seconds:", err)
      time.Sleep(time.Duration(2) * time.Second)
    }
  }

  db.AutoMigrate(
    &User{},
    &CredRecord{},
    &Peer{},
    &SentMessage{},
  )

  db.Model(&CredRecord{}).AddUniqueIndex("idx_cred_hash", "cred_hash")
  db.Model(&CredRecord{}).AddIndex("idx_statement_hash", "statement_hash")
  db.Model(&SentMessage{}).AddUniqueIndex("idx_message_hash", "message_hash")
  db.Model(&User{}).AddUniqueIndex("idx_fingerprint", "fingerprint")
}

func DB() gorm.DB {
  return db
}
