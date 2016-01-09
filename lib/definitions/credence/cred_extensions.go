package credence

import (
  "encoding/hex"
  "github.com/golang/protobuf/proto"
  "github.com/spacemonkeygo/openssl"
  "github.com/jphastings/credence/lib/models"
)

func (cred *Cred) StatementHash() string {
  statementCred := &Cred{
    Statement: cred.Statement,
  }

  statementCredBytes, err := proto.Marshal(statementCred)
  if err != nil {
    // TODO: Deal with error
    panic(err)
  }

  hash, _ := openssl.SHA1(statementCredBytes)
  hashBytes := []byte{}
  for _, b := range hash {
    hashBytes = append(hashBytes, b)
  }

  return hex.EncodeToString(hashBytes)
}

func (cred *Cred) SetSignature() error {
  sigCred := &Cred{}
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

func (cred *Cred) DetectAuthor() model.User {

}

