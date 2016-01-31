package helpers

import (
  "github.com/golang/protobuf/proto"
  "github.com/jphastings/credence/lib/models"
  "github.com/jphastings/credence/lib/definitions/credence"
)

func AssertIdentity(user *models.User) *credence.IdentityAssertion {
  me := models.Me()

  identityAssertion := &credence.IdentityAssertion{
    PublicKey: user.PublicKey,
    Name: user.Name,
    IdentityUri: user.IdentityUri,
  }

  if HasPrivateKey() {
    identityAssertion.Fingerprint = me.Fingerprint

    bytes, err := proto.Marshal(identityAssertion)
    if err != nil {
      panic(err)
    }
    identityAssertion.Signature = SignBytes(bytes)
  }

  return identityAssertion
}
