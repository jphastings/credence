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

// Returns the author of the cred, if it can be determined.
// Will return an error if the public key for the author_fingerprint
// doesn't match the signature.
func DetectAuthor(cred *credence.Cred) (models.User, error) {
  author := models.User{}
  db := models.DB()

  if (cred.AuthorFingerprint == nil) {
    users := []models.User{}
    db.Where("public_key IS NOT NULL").Find(&users)

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
  } else {
    db.Where("fingerprint = ?", hex.EncodeToString(cred.AuthorFingerprint)).First(&author)

    if !db.NewRecord(author) {
      sigCredByte := SignableCredBytes(cred)
      publicKey, err := openssl.LoadPublicKeyFromPEM(author.PublicKey)
      if err == nil {
        verifyErr := publicKey.VerifyPKCS1v15(openssl.SHA256_Method, sigCredByte, cred.Signature)
        if verifyErr != nil {
          return models.User{}, verifyErr
        }
      }
      // TODO: What if we can't load the public key?
    }
  }

  return author, nil
}

func CredUri(cred *credence.Cred) string {
  credBytes, _ := proto.Marshal(cred)
  b64 := base64.URLEncoding.EncodeToString(credBytes)
  // TODO: Figure out why RawURLEncoding doesn't work…
  b64NoPadding := strings.Replace(b64, "=", "", -1)
  // Use the 'official' Credence server for now, but switch to
  // a web+credence: protocol in the future, for decentralised goodness.
  return fmt.Sprintf("http://getcredence.net/creds/info?cred=%s", b64NoPadding)
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
