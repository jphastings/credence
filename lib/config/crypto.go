package config

import (
  "os/user"
  "io/ioutil"
  "path/filepath"
  "github.com/spacemonkeygo/openssl"
)

var privateKey openssl.PrivateKey
var privateKeyLoaded bool = false

func PemPath() string {
  usr, _ := user.Current()
  return filepath.Join(usr.HomeDir, ".credence", "id_rsa")
}

func PrivateKey() (openssl.PrivateKey, error) {
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

  return privateKey
}
