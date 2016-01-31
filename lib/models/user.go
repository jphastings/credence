package models

import (
  "time"
)

type User struct {
  ID uint `gorm:"primary_key"`
  Name string
  // The public key fingerprint
  Fingerprint []byte
  // This user's public key
  PublicKey []byte
  // An identifying URI
  IdentityUri string
  Weight int `sql:"default 1"`

  CreatedAt time.Time
  UpdatedAt time.Time
}

func Me() User {
  user := &User{}
  db.FirstOrCreate(&user, User{ID: 1})
  return *user
}