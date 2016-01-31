package helpers

import (
  "os"
  "log"
  "io/ioutil"
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

  SavePublicKeyToDB(privateKey)

  return privateKey
}

func SavePublicKeyToDB(privateKey openssl.PrivateKey) {
  publicDerBlock, err := privateKey.MarshalPKIXPublicKeyDER()
  if err != nil {
    panic(err)
  }

  fingerprint, err := openssl.SHA256(publicDerBlock)
  if err != nil {
    panic(err)
  }

  me := models.Me()
  me.PublicKey = publicDerBlock
  me.Fingerprint = fingerprint[:]
  db := models.DB()
  db.Save(&me)
  log.Print("Stored self public key in user DB")
}

func SavePublicKeyIfNeccessary() {
  privateKey, err := LoadPrivateKey()
  if err != nil {
    return
  }

  me := models.Me()
  if me.Fingerprint == nil {
    SavePublicKeyToDB(privateKey)
  }
}

func HasPrivateKey() bool {
  _, err := os.Stat(PemPath())
  return err == nil
}

func SignBytes(bytes []byte) []byte {
  privateKey, err := LoadPrivateKey()
  if err != nil {
    panic(err)
  }

  sig, err := privateKey.SignPKCS1v15(openssl.SHA256_Method, bytes)
  if err != nil {
    panic(err)
  }

  return sig
}