package helpers

import (
  "encoding/hex"
  "github.com/golang/protobuf/proto"
  "github.com/spacemonkeygo/openssl"
  "github.com/jphastings/credence/lib/config"
  "github.com/jphastings/credence/lib/models"
  "github.com/jphastings/credence/lib/definitions/credence"
)

func StatementHash(cred *credence.Cred) string {
  statementCred := &credence.Cred{
    Statement: cred.Statement,
  }

  statementCredBytes, err := proto.Marshal(statementCred)
  if err != nil {
    panic(err)
  }

  hash, err := openssl.SHA1(statementCredBytes)
  if err != nil {
    panic(err)
  }

  return hex.EncodeToString(hash[:])
}

func SetSignature(cred *credence.Cred) error {
  sigCred := &credence.Cred{}
  sigCred = cred
  sigCred.Signature = []byte{}

  sigCredBytes, err := proto.Marshal(sigCred)
  if err != nil {
    return err
  }

  privateKey, err := config.PrivateKey()
  if err != nil {
    return err
  }

  cred.Signature, err = privateKey.SignPKCS1v15(openssl.SHA256_Method, sigCredBytes)
  if err != nil {
    return err
  }

  return nil
}

func DetectAuthor(cred *credence.Cred) models.User {
  return models.Me()
}

