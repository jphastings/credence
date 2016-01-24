package helpers

import (
  "fmt"
  "strings"
  "encoding/hex"
  "encoding/base64"
  "github.com/golang/protobuf/proto"
  "github.com/spacemonkeygo/openssl"
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
  sigCredBytes := SignableCredBytes(cred)

  privateKey, err := LoadPrivateKey()
  if err != nil {
    return err
  }

  cred.Signature, err = privateKey.SignPKCS1v15(openssl.SHA256_Method, sigCredBytes)
  if err != nil {
    return err
  }

  return nil
}

func SignableCredBytes(cred *credence.Cred) []byte {
  sigCred := &credence.Cred{}
  *sigCred = *cred
  sigCred.Signature = []byte{}

  sigCredBytes, err := proto.Marshal(sigCred)
  if err != nil {
    panic(err)
  }

  return sigCredBytes
}

func DetectAuthor(cred *credence.Cred) models.User {
  db := models.DB()
  users := []models.User{}
  db.Where("public_key IS NOT NULL").Find(&users)

  author := models.User{}

  sigCredByte := SignableCredBytes(cred)

  for _, user := range users {
    publicKey, err := openssl.LoadPublicKeyFromPEM(user.PublicKey)
    if err == nil {
      verifyErr := publicKey.VerifyPKCS1v15(openssl.SHA256_Method, sigCredByte, cred.Signature)
      if verifyErr == nil {
        author = user
        break
      }
    }
  }

  return author
}

func CredUri(cred *credence.Cred) string {
  credBytes, _ := proto.Marshal(cred)
  b64 := base64.URLEncoding.EncodeToString(credBytes)
  // TODO: Figure out why RawURLEncoding doesn't work…
  b64NoPadding := strings.Replace(b64, "=", "", -1)
  return fmt.Sprintf("credence:creds/info?cred=%s", b64NoPadding)
}

func CredFromBase64(b64 string) (*credence.Cred, error) {
  cred := &credence.Cred{}

  credBytes, err := base64.URLEncoding.DecodeString(b64)
  if err != nil {
    return cred, err
  }
  
  err = proto.Unmarshal(credBytes, cred)
  if err != nil {
    return cred, err
  }

  return cred, nil
}

