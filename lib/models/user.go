package models

type User struct {
  ID int `sql:"AUTO_INCREMENT"`
  Name string
  // The public key fingerprint
  Fingerprint string
  // This user's public key
  PublicKey []byte
  // An identifying URI
  IdentityUri string
  Weight int
}

func Me() User {
  user := &User{}
  db.First(&user, 1)
  return *user
}