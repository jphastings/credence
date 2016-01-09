package helpers

import (
  "io/ioutil"
  "encoding/hex"
  "github.com/spacemonkeygo/openssl"
  "github.com/jphastings/credence/lib/models"
  "github.com/jphastings/credence/lib/config"
)

var privateKey openssl.PrivateKey
var privateKeyLoaded bool = false

func PemPath() string {
  return config.ConfigFile("id_rsa")
}

func LoadPrivateKey() (openssl.PrivateKey, error) {
  var err error

  if !privateKeyLoaded {
    pemBytes, err := ioutil.ReadFile(PemPath())

    if err != nil {
      // File doesn't exist
      return privateKey, err
    }

    privateKey, err = openssl.LoadPrivateKeyFromPEM(pemBytes)
    privateKeyLoaded = true
  }

  return privateKey, err
}

func CreatePrivateKey() openssl.PrivateKey {
  privateKey, _ := openssl.GenerateRSAKey(2048)
  pemBlock, err := privateKey.MarshalPKCS1PrivateKeyPEM()
  if err != nil {
    panic(err)
  }

  ioutil.WriteFile(PemPath(), pemBlock, 0600)

  publicPemBlock, err := privateKey.MarshalPKIXPublicKeyPEM()
  if err != nil {
    panic(err)
  }

  fingerprint, err := openssl.SHA1(publicPemBlock)
  if err != nil {
    panic(err)
  }

  me := models.Me()
  me.PublicKey = publicPemBlock
  me.Fingerprint = hex.EncodeToString(fingerprint[:])
  db := models.DB()
  db.Save(&me)

  return privateKey
}
