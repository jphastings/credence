package viewModels

import (
  "reflect"
  "github.com/jphastings/credence/lib/models"
  "github.com/jphastings/credence/lib/definitions/credence"
)

type DisplayedUri struct {
  String string
  Url string
}

type Props interface {}

func Retrieve(data interface{}) (string, Props) {
  switch reflect.TypeOf(data).Elem() {
  case reflect.TypeOf(credence.IdentityAssertion{}):
    return "identity_assertion", RetrieveIdentityAssertion(data.(*credence.IdentityAssertion))
  case reflect.TypeOf(models.CredRecord{}):
    return "cred_record", RetrieveCredRecord(data.(*models.CredRecord))
  default:
    panic(data)
  }
  return "", nil
}